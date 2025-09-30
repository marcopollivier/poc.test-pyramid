.PHONY: up down build run test test-unit test-integration

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

test-unit:
	go test ./service/... -v

test-integration: up
	sleep 10
	go test ./integration_test/... -v

clean:
	docker compose down -v
	rm -f bin/app
