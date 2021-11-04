.PHONY: all

current_dir = $(shell pwd)

all: gettools generate format build

gettools:
	go install golang.org/x/tools/cmd/goimports@v0.1.5
	go mod tidy

generate:
	go generate ./...

format:
	goimports -w .
	go mod tidy

download:
	go mod download

build: download
	go build ./cmd/generateblock/
