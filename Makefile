SHELL := /bin/bash

.PHONY: all check format vet lint build release clean test coverage

VERSION=`git describe`

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check      to format, vet and lint "
	@echo "  build      to create bin directory and build go-mod-redirect"
	@echo "  release    to release go-mod-redirect"
	@echo "  clean      to clean build and test files"
	@echo "  test       to run test"
	@echo "  coverage   to test with coverage"

check: format vet lint

format:
	@echo "go fmt"
	@go fmt ./...
	@echo "ok"

vet:
	@echo "go vet"
	@go vet ./...
	@echo "ok"

lint:
	@echo "golint"
	@golint ./...
	@echo "ok"

build: check
	@echo "build go-mod-redirect"
	@mkdir -p ./bin
	@go build -tags netgo -o ./bin/go-mod-redirect ./cmd/go-mod-redirect
	@echo "ok"

release:
	@echo "release go-mod-redirect"
	@-rm ./release/*
	@mkdir -p ./release

	@echo "build for linux"
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/go-mod-redirect_${VERSION}_linux_amd64 ./cmd/go-mod-redirect
	@tar -C ./bin/linux/ -czf ./release/go-mod-redirect_${VERSION}_linux_amd64.tar.gz go-mod-redirect_${VERSION}_linux_amd64

	@echo "ok"

clean:
	@rm -rf ./bin
	@rm -rf ./release
	@rm -rf ./coverage

test:
	@echo "run test"
	@go test -v ./...
	@echo "ok"

coverage:
	@echo "run test with coverage"
	@mkdir -p coverage
	@go test -v -cover -coverprofile="coverage/profile.out" ./...
	@go tool cover -html="coverage/profile.out" -o "coverage/profile.html"
	@echo "ok"