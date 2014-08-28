GOPATH = $(shell pwd)
DISTDIR = ./dist
EXEC = xconf-go-svc
VERSION = $(shell cat VERSION)
REVISION = $(shell git log --pretty=format:'%h' -n 1)

.deps:
	GOPATH=$(GOPATH) go get -d

.build: .deps
	GOPATH=$(GOPATH) go build -o $(EXEC)

.dist-dir:
	mkdir -p $(DISTDIR)

default: .build

clean:
	rm -rf $(DISTDIR)
	rm -f $(EXEC)

tar: .build .dist-dir
	mkdir -p $(DISTDIR)/usr/bin
	mv $(EXEC) $(DISTDIR)/usr/bin
	cp -R etc/ $(DISTDIR)/
	tar -czO -C $(DISTDIR) usr etc > $(DISTDIR)/$(EXEC).tar.gz

package: .build .dist-dir
	fpm -s dir \
		-t rpm \
		-n $(EXEC) \
		-v $(VERSION)-$(REVISION) \
		-p dist/$(EXEC).rpm \
		--no-depends \
		$(EXEC)=/usr/bin/$(EXEC) \
		etc/init.d/$(EXEC)=/etc/init.d/$(EXEC) \
		etc/default/$(EXEC)=/etc/default/$(EXEC)
