.PHONY: up down build run test

up:
	docker compose up -d

down:
	docker compose down

build:
	go build -o bin/app .

run: up
	sleep 5
	go run main.go

test:
	go test ./... -v

clean:
	docker compose down -v
	rm -f bin/app
