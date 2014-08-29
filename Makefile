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

.service-script: .dist-dir
	echo "service $(EXEC) restart" >> $(DISTDIR)/after-install
	echo "service $(EXEC) stop" >> $(DISTDIR)/after-remove
	chmod +x $(DISTDIR)/after-install
	chmod +x $(DISTDIR)/after-remove

default: .build

clean:
	rm -rf $(DISTDIR)
	rm -f $(EXEC)

cp: .build
	cp $(EXEC) /usr/bin
	cp ./etc/init.d/$(EXEC) /etc/init.d
	cp ./etc/default/$(EXEC) /etc/default

rsync: .build
	$(eval NAME := $(shell read -p "Citizen... what is your name? " name && echo $$name))
	rsync --rsync-path="sudo rsync" -a $(EXEC) vagrant@$(NAME)-web-service.local:/usr/bin
	rsync --rsync-path="sudo rsync" -a etc/init.d/$(EXEC) vagrant@$(NAME)-web-service.local:/etc/init.d
	rsync --rsync-path="sudo rsync" -a etc/default/$(EXEC) vagrant@$(NAME)-web-service.local:/etc/default

tar: .build .dist-dir
	mkdir -p $(DISTDIR)/usr/bin
	mv $(EXEC) $(DISTDIR)/usr/bin
	cp -R etc/ $(DISTDIR)/
	tar -czO -C $(DISTDIR) usr etc > $(DISTDIR)/$(EXEC).tar.gz

rpm: .build .dist-dir
	fpm -s dir \
		-f \
		-t rpm \
		-n $(EXEC) \
		-v $(VERSION)-$(REVISION) \
		-p dist/$(EXEC).rpm \
		--no-depends \
		$(EXEC)=/usr/bin/$(EXEC) \
		etc/init.d/$(EXEC)=/etc/init.d/$(EXEC) \
		etc/default/$(EXEC)=/etc/default/$(EXEC)

rpm-with-script: .build .dist-dir .service-script
	fpm -s dir \
		-f \
		-t rpm \
		-n $(EXEC) \
		-v $(VERSION)-$(REVISION) \
		-p dist/$(EXEC).rpm \
		--after-install $(DISTDIR)/after-install \
		--after-remove $(DISTDIR)/after-remove \
		--no-depends \
		$(EXEC)=/usr/bin/$(EXEC) \
		etc/init.d/$(EXEC)=/etc/init.d/$(EXEC) \
		etc/default/$(EXEC)=/etc/default/$(EXEC)

rpm-with-deps: .build .dist-dir .service-script
	fpm -s dir \
		-f \
		-t rpm \
		-n $(EXEC) \
		-v $(VERSION)-$(REVISION) \
		-p dist/$(EXEC).rpm \
		--after-install $(DISTDIR)/after-install \
		--after-remove $(DISTDIR)/after-remove \
		-d mysql-libs -d redis -d python \
		$(EXEC)=/usr/bin/$(EXEC) \
		etc/init.d/$(EXEC)=/etc/init.d/$(EXEC) \
		etc/default/$(EXEC)=/etc/default/$(EXEC)

package: rpm-with-script
