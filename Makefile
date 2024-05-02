.DEFAULT_GOAL := build

.PHONY:clean
clean:
	rm -rf output

.PHONY:fmt
fmt:
	go fmt ./...

.PHONY:vet
vet: fmt
	go vet ./...

.PHONY:vet
build: clean vet
	go build -o output/ ./...

.PHONY:docker
docker:
	docker compose -f deployments/docker-compose.yml down
	docker compose -f deployments/docker-compose.yml build
	docker compose -f deployments/docker-compose.yml up --watch