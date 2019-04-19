
SRCPATH := github.com/amaltbie/sparky
SOURCE := $(shell find $(PWD) -name '*.go')
USERSTR := $(shell id -u):$(shell id -g)

.docker: docker/Dockerfile
	docker build -t sparky-build docker/
	touch .docker

.build/sprk: $(SOURCE) .docker
	mkdir -p .cache
	mkdir -p .build
	mkdir -p .gopath
	docker run \
		-v $(PWD):/go/src/$(SRCPATH) \
		-u $(USERSTR) \
		-v $(PWD)/.cache:/.cache \
		-v $(PWD)/.gopath:/gopath \
		-e GOOS=darwin \
		-e GO111MODULE=on \
		-e GOPATH=/gopath \
		-w /go/src/$(SRCPATH) sparky-build \
		go build -o /go/src/$(SRCPATH)/.build/sprk

build: .build/sprk

clean:
	rm .build/sprk
