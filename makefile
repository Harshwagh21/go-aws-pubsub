.PHONY: up down run logs build clean

up: 
	docker-compose up -d

down:
	docker-compose down

run:
	go run ./cmd/main.go

logs:
	docker logs localstack

build:
	go build ./...

clean:
	docker-compose down
	go clean