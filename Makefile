.PHONY: build clean tool help

all: build

build:
	go build .

tool:
	go tool vet . | grep -v vendor; true
	gofmt -w .

clean:
	rm -rf go-gin-example
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make clean: remove object files and cached files"