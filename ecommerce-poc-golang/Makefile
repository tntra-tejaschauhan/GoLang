.PHONY: help setup start stop clean test deps build run-all

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup: ## Initial setup - start databases and install dependencies
	@echo "Starting databases..."
	docker-compose up -d
	@echo "Waiting for databases to be ready..."
	sleep 5
	@echo "Installing dependencies..."
	@cd gateway-app && go mod download
	@cd product-app && go mod download
	@cd cart-app && go mod download
	@cd invoice-app && go mod download
	@cd payment-service && go mod download
	@cd management-service && go mod download
	@echo "Setup complete!"

deps: ## Download dependencies for all services
	@echo "Downloading dependencies..."
	@cd gateway-app && go mod download && go mod tidy
	@cd product-app && go mod download && go mod tidy
	@cd cart-app && go mod download && go mod tidy
	@cd invoice-app && go mod download && go mod tidy
	@cd payment-service && go mod download && go mod tidy
	@cd management-service && go mod download && go mod tidy
	@echo "Dependencies updated!"

db-up: ## Start PostgreSQL databases
	docker-compose up -d
	@echo "Waiting for databases to be ready..."
	sleep 5

db-down: ## Stop PostgreSQL databases
	docker-compose down

db-reset: ## Reset all databases (WARNING: destroys all data)
	docker-compose down -v
	docker-compose up -d
	@echo "Waiting for databases to be ready..."
	sleep 5

run-gateway: ## Run Gateway service
	@cd gateway-app && go run main.go

run-product: ## Run Product service
	@cd product-app && go run main.go

run-cart: ## Run Cart service
	@cd cart-app && go run main.go

run-invoice: ## Run Invoice service
	@cd invoice-app && go run main.go

run-payment: ## Run Payment service
	@cd payment-service && go run main.go

run-management: ## Run Management service
	@cd management-service && go run main.go

build: ## Build all services
	@echo "Building services..."
	@cd gateway-app && go build -o ../bin/gateway main.go
	@cd product-app && go build -o ../bin/product main.go
	@cd cart-app && go build -o ../bin/cart main.go
	@cd invoice-app && go build -o ../bin/invoice main.go
	@cd payment-service && go build -o ../bin/payment main.go
	@cd management-service && go build -o ../bin/management main.go
	@echo "Build complete! Binaries in ./bin/"

clean: ## Clean built binaries
	rm -rf bin/
	@echo "Cleaned!"

test-health: ## Test health endpoints of all services
	@echo "Testing health endpoints..."
	@curl -s http://localhost:8080/health || echo "Gateway: NOT RUNNING"
	@curl -s http://localhost:8081/health || echo "Product: NOT RUNNING"
	@curl -s http://localhost:8082/health || echo "Cart: NOT RUNNING"
	@curl -s http://localhost:8083/health || echo "Invoice: NOT RUNNING"
	@curl -s http://localhost:8084/health || echo "Payment: NOT RUNNING"
	@curl -s http://localhost:8085/health || echo "Management: NOT RUNNING"

test-flow: ## Run complete test flow
	@echo "Creating product..."
	@curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d '{"name":"Test Laptop","price":899.99}'
	@echo "\nAdding to cart..."
	@curl -X POST http://localhost:8080/cart/items -H "Content-Type: application/json" -H "X-User-Id: user-1" -d '{"product_id":1,"quantity":1,"price":899.99}'
	@echo "\nChecking out..."
	@curl -X POST http://localhost:8080/cart/checkout -H "X-User-Id: user-1"
	@echo "\nTest complete!"

stop: ## Stop all running Go processes
	@echo "Stopping all services..."
	@pkill -f "go run main.go" || echo "No services running"
	@echo "Services stopped!"

logs-db: ## Show database logs
	docker-compose logs -f

ps: ## Show running processes
	@echo "Docker containers:"
	@docker-compose ps
	@echo "\nGo processes:"
	@ps aux | grep "go run" | grep -v grep || echo "No Go services running"

fmt: ## Format Go code
	@cd gateway-app && go fmt ./...
	@cd product-app && go fmt ./...
	@cd cart-app && go fmt ./...
	@cd invoice-app && go fmt ./...
	@cd payment-service && go fmt ./...
	@cd management-service && go fmt ./...
	@echo "Code formatted!"

lint: ## Run linters (requires golangci-lint)
	@cd gateway-app && golangci-lint run
	@cd product-app && golangci-lint run
	@cd cart-app && golangci-lint run
	@cd invoice-app && golangci-lint run
	@cd payment-service && golangci-lint run
	@cd management-service && golangci-lint run
