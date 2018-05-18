BINPATH := bin
SERVER_BIN_NAME := math-server
CLIENT_BIN_NAME := math-client

generate: pb/service.proto
	go generate ./...

test:
	go test -v ./...

build: build-server build-client

build-server:
	go build -o $(BINPATH)/$(SERVER_BIN_NAME) cmd/server/*.go

build-client:
	go build -o $(BINPATH)/$(CLIENT_BIN_NAME) cmd/client/*.go

clean:
	rm -rf $(BINPATH)

.PHONY: test clean

