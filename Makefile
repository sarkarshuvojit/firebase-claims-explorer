default:
	@echo "Cmds [test, build, run]"

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
HASH := $(shell git rev-parse --short HEAD)

test:
	@go test ./...

build:
	@go build -o .bin/asc-$(BRANCH)-$(HASH)

run:
	@go run main.go explore --config .test/credentials.json

