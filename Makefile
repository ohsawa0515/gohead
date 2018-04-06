.PHONY: build test

APP := head
VERSION := "0.0.1"
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	-X 'main.revision=$(REVISION)'

test:
	go test -v ./...

build:
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(PWD)/bin/darwin_amd64/$(APP)
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(PWD)/bin/linux_amd64/$(APP)