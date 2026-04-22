.PHONY: run build test fix docker-up docker-down swagger

# Variables
BINARY_NAME=myservice
MAIN_PATH=cmd/api/main.go

run:
	go run $(MAIN_PATH)

build:
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

test:
	go test ./... -v

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

tidy:
	go mod tidy

fmt:
	go fmt ./...

swagger:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/api/main.go -o docs

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down
