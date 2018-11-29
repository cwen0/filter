GO=GO15VENDOREXPERIMENT="1" CGO_ENABLED=0 GO111MODULE=on go
GOBUILD=$(GO) build
GOTEST=GO15VENDOREXPERIMENT="1" CGO_ENABLED=0 GO111MODULE=on go test

default: build

build:
	$(GOBUILD) -o bin/filter cmd/filter/*.go

test:
	$(GOTEST) ./...
