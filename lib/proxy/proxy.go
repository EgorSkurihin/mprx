package proxy

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type ProxyParams struct {
	ProxyAddr    string
	ResourceAddr string
}

type Proxy struct {
	proxyAddr    string
	resourceAddr string
	connListener net.Listener
	stopCh       chan struct{}
	wg           sync.WaitGroup
}

func NewProxy(params ProxyParams) (*Proxy, error) {
	connListener, err := net.Listen("tcp", params.ProxyAddr)
	if err != nil {
		return nil, fmt.Errorf("can`t start proxy server: %s", err.Error())
	}
	p := &Proxy{
		proxyAddr:    params.ProxyAddr,
		resourceAddr: params.ResourceAddr,
		connListener: connListener,
		stopCh:       make(chan struct{}),
	}
	return p, nil
}

func (p *Proxy) Start() {
	for {
		conn, err := p.connListener.Accept()
		if err != nil {
			select {
			case <-p.stopCh:
				return
			default:
				log.Printf("error, failed to receive connection: %v", err)
			}
		} else {
			p.wg.Add(1)
			go func() {
				defer p.wg.Done()
				p.handleConnection(conn)
			}()
		}
	}
}

func (p *Proxy) Stop() {
	close(p.stopCh)
	p.connListener.Close()
	p.wg.Wait()
}

func (p *Proxy) handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("handle connection from %s", conn.RemoteAddr().String())
	resource, err := net.Dial("tcp", p.resourceAddr)
	if err != nil {
		log.Printf("error, failed to connect to mongodb: %v", err)
		return
	}

	go mongoToApp(conn, resource)
	appToMongo(conn, resource)
}

func mongoToApp(app net.Conn, mongo net.Conn) {
	for {
		pkt, err := proxyPacket(mongo, app)
		if err != nil {
			break
		}
		log.Printf("msg from server to client: %s", pkt)
	}
}

func appToMongo(app net.Conn, mongo net.Conn) {
	for {
		pkt, err := proxyPacket(app, mongo)
		if err != nil {
			break
		}
		log.Printf("msg from client to server: %s", pkt)
	}
}
