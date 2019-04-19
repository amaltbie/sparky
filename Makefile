
SRCPATH := github.com/amaltbie/sparky
SOURCE := $(shell find $(PWD) -name '*.go')
USERSTR := $(shell id -u):$(shell id -g)
OSTYPE := $(shell uname | tr '[:upper:]' '[:lower:]')

ifndef DESTDIR
	DESTDIR := /usr/local
endif
ifndef VERSION
	VERSION := local-$(shell git rev-parse HEAD | cut -c-8)
endif

.docker: docker/Dockerfile
	docker build -t sparky-build docker/
	touch .docker

.cache:
	mkdir -p .cache

.gopath:
	mkdir -p .gopath

.build:
	mkdir -p .build

.build/sprk-darwin: $(SOURCE) .docker | .cache .gopath .build
	docker run \
		-v $(PWD):/go/src/$(SRCPATH) \
		-u $(USERSTR) \
		-v $(PWD)/.cache:/.cache \
		-v $(PWD)/.gopath:/gopath \
		-e GOOS=darwin \
		-e GO111MODULE=on \
		-e GOPATH=/gopath \
		-w /go/src/$(SRCPATH) sparky-build \
		go build \
		-o /go/src/$(SRCPATH)/.build/sprk-darwin \
		-ldflags "-X github.com/amaltbie/sparky/cmd.Version=${VERSION}"

.build/sprk-linux: $(SOURCE) .docker | .cache .gopath .build
	docker run \
		-v $(PWD):/go/src/$(SRCPATH) \
		-u $(USERSTR) \
		-v $(PWD)/.cache:/.cache \
		-v $(PWD)/.gopath:/gopath \
		-e GOOS=linux \
		-e GO111MODULE=on \
		-e GOPATH=/gopath \
		-w /go/src/$(SRCPATH) sparky-build \
		go build \
		-o /go/src/$(SRCPATH)/.build/sprk-linux \
		-ldflags "-X github.com/amaltbie/sparky/cmd.Version=${VERSION}"

build: .build/sprk-darwin .build/sprk-linux

clean:
	rm -rf .build

clobber: clean
	rm -rf .gopath
	rm -rf .cache

install:
	mkdir -p $(DESTDIR)/bin
	install .build/sprk-$(OSTYPE) $(DESTDIR)/bin/sprk

.DEFAULT_GOAL := build
