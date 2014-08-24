GOPATH = $(shell pwd)

.deps:
	GOPATH=$(GOPATH) go get -d

default: .deps
	GOPATH=$(GOPATH) go build
