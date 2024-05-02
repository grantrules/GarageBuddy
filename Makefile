.DEFAULT_GOAL := build
.PHONY:fmt vet build clean

clean:
	rm -rf output
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: clean vet
	go build -o output/ ./...