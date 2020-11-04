GOPATH:=$(shell go env GOPATH)

.PHONY: build compile clean

build: clean
	@go build -o dist/engine *.go

compile: clean
	@go build -tags static -ldflags "-s -w" -o dist/engine *.go

clean:
	@rm -rf dist/engine