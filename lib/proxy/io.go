package proxy

import (
	"io"
	"net"
)

const bsonHeaderLength = 4

// ProxyPacket объединяет методы ReadPacket и WritePacket
func proxyPacket(src, dst net.Conn) ([]byte, error) {
	pkt, err := readPacket(src)
	if err != nil {
		return nil, err
	}

	_, err = writePacket(pkt, dst)
	if err != nil {
		return nil, err
	}

	return pkt, nil
}

// ReadPacket читает данные из conn, возвращая готовый пакет
func readPacket(conn net.Conn) ([]byte, error) {
	header := make([]byte, bsonHeaderLength)

	if _, err := io.ReadFull(conn, header); err == io.EOF {
		return nil, io.ErrUnexpectedEOF
	} else if err != nil {
		return nil, err
	}

	bodyLength := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)

	body := make([]byte, bodyLength-bsonHeaderLength)

	n, err := io.ReadFull(conn, body)
	if err == io.EOF {
		return nil, io.ErrUnexpectedEOF
	} else if err != nil {
		return nil, err
	}

	return append(header, body[0:n]...), nil
}

// WritePacket пишет пакет, полученный из метода ReadPacket, в conn
func writePacket(pkt []byte, conn net.Conn) (int, error) {
	n, err := conn.Write(pkt)
	if err != nil {
		return 0, err
	}

	return n, nil
}
