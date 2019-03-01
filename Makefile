BINARY=bastion
BIN_PATH=bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all: build

build:
	$(GOBUILD)

release: clean generate build

run:
	$(GOBUILD) -o $(BIN_PATH)/$(BINARY) -v
	env $(BIN_PATH)/$(BINARY)

clean:
	$(GOCLEAN)
	rm -f $(BIN_PATH)/*

generate:
	$(GOGEN)

deps:
	$(GOGET) -d ./...

env:
	cp .env.sample .env