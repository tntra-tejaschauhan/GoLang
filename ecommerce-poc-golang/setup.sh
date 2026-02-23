#!/bin/bash

# E-Commerce POC - Setup and Run Script

set -e

echo "=========================================="
echo "E-Commerce POC - Golang Setup"
echo "=========================================="

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker first."
    exit 1
fi

echo ""
echo "1️⃣  Starting PostgreSQL databases..."
# docker-compose up -d

# echo "⏳ Waiting for databases to be ready..."
# sleep 5

echo ""
echo "2️⃣  Installing dependencies for all services..."

services=("gateway-app" "product-app" "cart-app" "invoice-app" "payment-service" "management-service")

for service in "${services[@]}"; do
    echo "   📦 Installing dependencies for $service..."
    cd "$service"
    go mod download
    cd ..
done

echo ""
echo "✅ Setup complete!"
echo ""
echo "=========================================="
echo "To run the services:"
echo "=========================================="
echo ""
echo "Open 6 separate terminal windows and run:"
echo ""
echo "Terminal 1: cd gateway-app && go run main.go"
echo "Terminal 2: cd product-app && go run main.go"
echo "Terminal 3: cd cart-app && go run main.go"
echo "Terminal 4: cd invoice-app && go run main.go"
echo "Terminal 5: cd payment-service && go run main.go"
echo "Terminal 6: cd management-service && go run main.go"
echo ""
echo "=========================================="
echo "Service URLs:"
echo "=========================================="
echo "Gateway:    http://localhost:8080"
echo "Product:    http://localhost:8081"
echo "Cart:       http://localhost:8082"
echo "Invoice:    http://localhost:8083"
echo "Payment:    http://localhost:8084"
echo "Management: http://localhost:8085"
echo ""
echo "See API_EXAMPLES.md for testing examples"
echo "=========================================="
