package cmdproxy

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/EgorSkurihin/mprx/lib/proxy"
)

type proxyFacade struct {
	params *Proxy
	proxy  *proxy.Proxy

	term chan os.Signal
}

func newProxyFacade(params *Proxy) (*proxyFacade, error) {

	facade := &proxyFacade{
		params: params,
	}
	if err := facade.createProxyServer(); err != nil {
		return nil, err
	}
	return facade, nil
}

func (facade *proxyFacade) start() error {
	//defer facade.Stop()

	facade.term = make(chan os.Signal, 1)
	signal.Notify(facade.term, syscall.SIGTERM, syscall.SIGINT)

	go facade.proxy.Start()
	log.Println("info, proxy server successfully started...")
	facade.proxyServerProcess()

	return nil
}

func (facade *proxyFacade) Stop() {
	signal.Stop(facade.term)
	close(facade.term)
	log.Println("info, proxy server stoped")
}

func (facade *proxyFacade) createProxyServer() error {
	p, err := proxy.NewProxy(proxy.ProxyParams{
		ProxyAddr:    facade.params.ProxyAddr,
		ResourceAddr: facade.params.MongoDBAddr,
	})
	if err != nil {
		return fmt.Errorf("failed to create proxy server: %v", err)
	}
	facade.proxy = p
	return nil
}

func (facade *proxyFacade) proxyServerProcess() {
	if _, ok := <-facade.term; !ok {
		return
	}
	facade.Stop()
}
