.PHONY: help build run test clean docker-build docker-run install dev

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod verify

build: ## Build the application
	go build -o bin/url-shortener main.go

run: ## Run the application
	go run main.go

run-memory: ## Run with in-memory storage
	USE_IN_MEMORY=true go run main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-api: ## Test the API with test client
	@echo "Make sure the server is running (make run in another terminal)"
	@sleep 1
	go run test_client.go

lint: ## Run linter
	go fmt ./...
	go vet ./...

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f *.db
	rm -f coverage.out coverage.html

docker-build: ## Build Docker image
	docker build -t url-shortener:latest .

docker-run: ## Run Docker container
	docker run -p 8080:8080 -v $(PWD)/data:/data url-shortener:latest

docker-compose-up: ## Start with docker-compose
	docker-compose up -d

docker-compose-down: ## Stop docker-compose
	docker-compose down

docker-compose-logs: ## View docker-compose logs
	docker-compose logs -f

dev: ## Run in development mode with auto-reload (requires air)
	@which air > /dev/null || (echo "Installing air..." && go install github.com/cosmtrek/air@latest)
	air

.DEFAULT_GOAL := help

