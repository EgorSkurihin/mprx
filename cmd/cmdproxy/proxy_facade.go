package cmdproxy

import "fmt"

type proxyFacade struct {
}

func newProxyFacade(params *Proxy) (*proxyFacade, error) {
	return &proxyFacade{}, nil
}

func (facade *proxyFacade) start() error {
	defer facade.Stop()
	fmt.Println("server started")
	return nil
}

func (facade *proxyFacade) Stop() {
	fmt.Println("server stopped")
}
