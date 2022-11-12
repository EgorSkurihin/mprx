package cmdstart

type startFacade struct {
}

func newStartFacade(params *Start) (*startFacade, error) {
	return &startFacade{}, nil
}

func (facade *startFacade) startServer() error {
	return nil
}
