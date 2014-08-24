GOPATH = $(shell pwd)

.deps:
	GOPATH=$(GOPATH) go get -d

.build: .deps
	GOPATH=$(GOPATH) go build -o xconf-go-svc

default: .build

package: .build
	fpm -s dir -t rpm -n xconf-go-svc -v 1.0 --no-depends xconf-go-svc=/usr/bin/xconf-go-svc etc/init.d/xconf-go-svc=/etc/init.d/xconf-go-svc etc/default/xconf-go-svc=/etc/default/xconf-go-svc
