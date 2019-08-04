.PHONY: all clean build test

all:
	make build

build:
	go build -o ./bin/monsters ./cmd/monsters/main.go

clean:
	rm -f ./bin/*

test:
	go clean -testcache
	go test ./internal/...

