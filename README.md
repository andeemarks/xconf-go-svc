XConf User Service
==================

Background
----------

This service is nothing short of a cheap knock-off of [this basic Golang microservice](http://geeklit.blogspot.com.au/2013/08/gorest-golang-web-services-simple.html).  It has been re-produced here to provide a suitable starting point for a set of linked workshops comprising the Sydney 2014 version of XConf, an internal technical conference for ThoughtWorks employees.

Prerequisites
-------------

Because this service runs as a Golang application, you will need to have a Golang development environment installed on your machine and your `$GOPATH` environment variable set.  See [here](https://code.google.com/p/go-wiki/wiki/GOPATH) for instructions on how to do this.

* `go version` => `go version go1.2.2 darwin/amd64`

Dependencies
------------

* `go get github.com/emicklei/go-restful`
* `go get github.com/emicklei/go-restful/swagger`
* `go get bitbucket.org/kardianos/osext`
* `go get github.com/onsi/ginkgo/ginkgo`

Installation
------------

1. `cd $GOPATH/src`
1. `git clone https://github.com/andeemarks/xconf-go-svc.git` 

Building
--------

1. `cd $GOPATH/src/github.com/andeemarks/xconf-go-svc`
1. `go build`

Running
-------

1. `cd $GOPATH/src/github.com/andeemarks/xconf-go-svc`
1. `./xconf-go-svc`
1. `open http://localhost:8080/apidocs/` for a [Swagger](https://helloreverb.com/developers/swagger) browser showing the API.