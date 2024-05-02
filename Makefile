.DEFAULT_GOAL := build
.PHONY:fmt vet build clean

clean:
	rm -rf hello
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: clean vet
	go build