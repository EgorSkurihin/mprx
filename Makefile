.PHONY: all test build

all: test build

build:
	go build . 

test:
	GORACE="halt_on_error=1" go test -race ./...
