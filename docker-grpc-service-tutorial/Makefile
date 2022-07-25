SHELL = /bin/bash

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: proto
## proto: compiles .proto files
proto:
	@ docker run -v $(PWD):/defs namely/protoc-all -f proto/poetry.proto -l go -o . --go-source-relative

.PHONY: build
## build: builds server's binary
build:
	@ go build -a -installsuffix cgo -o main .

.PHONY: run
## run: runs the server
run: build
	@ ./main

.PHONY: test
## test: runs unit tests
test:
	@ go test -v ./...
