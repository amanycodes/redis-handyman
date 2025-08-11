APP=rhm
MAIN=./main.go

LDFLAGS=-s -w -X github.com/amanycodes/redis-handyman/cmd/rhm.version=$(shell git describe --tags --always --dirty 2>/dev/null || echo dev) \
        -X github.com/amanycodes/redis-handyman/cmd/rhm.commit=$(shell git rev-parse --short HEAD 2>/dev/null || echo none) \
        -X github.com/amanycodes/redis-handyman/cmd/rhm.date=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

.PHONY: build run test lint tidy fmt vet

build:
	go build -ldflags "$(LDFLAGS)" -o $(APP) $(MAIN)

run: build
	./$(APP) $(ARGS)

test:
	go test ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...
