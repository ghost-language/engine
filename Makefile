GOPATH:=$(shell go env GOPATH)

.PHONY: build clean

build: build-mac

build-mac: clean
	@go build -o dist/engine *.go

clean:
	@rm -rf dist/engine