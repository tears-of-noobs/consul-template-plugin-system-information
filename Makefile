.PHONY: all clean-all build cleand-deps deps ver test install

DATE := $(shell git log -1 --format="%cd" --date=short | sed s/-//g)
COUNT := $(shell git rev-list --count HEAD)
COMMIT := $(shell git rev-parse --short HEAD)

BINARYNAME := consul-template-plugin-system-information
VERSION := "${DATE}.${COUNT}_${COMMIT}"

LDFLAGS := "-X main.version=${VERSION}"


default: all 

all: clean-all deps build

ver:
	@echo ${VERSION}

clean-all: clean-deps
	@echo Clean builded binaries
	rm -rf .out/
	@echo Done

test:
	go test -v

build:
	@echo Build
	ln -s ${PWD}/vendor/ ${PWD}/vendor/src
	GOPATH="${PWD}/vendor" go build -v -o .out/${BINARYNAME} -ldflags ${LDFLAGS} *.go
	@echo Done

install: clean-all deps
	@echo Installing into ${GOPATH}/bin
	go install

clean-deps:
	@echo Clean dependencies
	rm -rf vendor/*

deps:
	@echo Fetch dependencies
	dep ensure
