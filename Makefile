.PHONY: all clean build test integration-test

all:
	make build

build:
	go build -o ./bin/monsters ./cmd/monsters/main.go

clean:
	rm -f ./bin/*

test:
	go clean -testcache
	go test ./internal/...

integration-test:
	make build
	go clean -testcache
	go test ./test/...

