.PHONY: clean install
.DEFAULT_GOAL := all

# Global vars
NAME		:= simpler
ORG 		:= simpler.com
GOARCH		:= amd64
VERSION   	?= $(shell cat VERSION | tr -d '\n')
GITCOMMIT	:= $(shell git rev-parse --short HEAD | tr -d '\n')
OS 			:= $(shell uname -s | tr -d '\n')
BUILD := $(shell date +%FT%T%z)

# Application vars
SIMPLER      = simpler

# Golang build vars
PKG=$(ORG)/$(NAME)
LDFLAGS="-X ${PKG}/pkg/config.Version=$(VERSION)\
	-X ${PKG}/pkg/config.Build=$(BUILD)\
	-X ${PKG}/pkg/config.GitCommit=$(GITCOMMIT)\
	-X ${PKG}/pkg/config.OS=$(OS)"

GO_ARGS := -ldflags ${LDFLAGS}

install:
	go build -o $(GOPATH)/bin/$(SIMPLER) -v $(VERBOSE_ARGS) $(GO_ARGS) ./cmd/simpler

build:
	go build -o build/$(SIMPLER) -v $(VERBOSE_ARGS) $(GO_ARGS) ./cmd/simpler

all: clean test lint install

test:
	go test ./...

lint:
	golint && golangci-lint run

clean:
	@rm -f $$GOPATH/bin/$(SIMPLER)
	@rm -f build/$(SIMPLER)
