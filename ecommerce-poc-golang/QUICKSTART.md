# Quick Start Guide

## Prerequisites

- **Go 1.21+** installed
- **Docker** and **Docker Compose** installed
- **Git** (for cloning)

## Installation Steps

### 1. Clone/Setup Project

```bash
cd ecommerce-poc-golang
```

### 2. Start Databases

```bash
docker-compose up -d
```

This starts 5 PostgreSQL containers:
- `postgres-product` on port **5432**
- `postgres-cart` on port **5433**
- `postgres-invoice` on port **5434**
- `postgres-payment` on port **5435**
- `postgres-management` on port **5436**

### 3. Install Dependencies

Run the setup script:
```bash
./setup.sh
```

Or manually for each service:
```bash
cd gateway-app && go mod download && cd ..
cd product-app && go mod download && cd ..
cd cart-app && go mod download && cd ..
cd invoice-app && go mod download && cd ..
cd payment-service && go mod download && cd ..
cd management-service && go mod download && cd ..
```

### 4. Start All Services

**Option A: Using separate terminals (Recommended for development)**

Open 6 terminal windows and run:

```bash
# Terminal 1 - Gateway (Port 8080)
cd gateway-app
go run main.go

# Terminal 2 - Product Service (Port 8081)
cd product-app
go run main.go

# Terminal 3 - Cart Service (Port 8082)
cd cart-app
go run main.go

# Terminal 4 - Invoice Service (Port 8083)
cd invoice-app
go run main.go

# Terminal 5 - Payment Service (Port 8084)
cd payment-service
go run main.go

# Terminal 6 - Management Service (Port 8085)
cd management-service
go run main.go
```

**Option B: Using background processes**

```bash
cd gateway-app && go run main.go &
cd product-app && go run main.go &
cd cart-app && go run main.go &
cd invoice-app && go run main.go &
cd payment-service && go run main.go &
cd management-service && go run main.go &
```

## Verify Installation

Check all services are running:

```bash
curl http://localhost:8080/health  # Gateway
curl http://localhost:8081/health  # Product
curl http://localhost:8082/health  # Cart
curl http://localhost:8083/health  # Invoice
curl http://localhost:8084/health  # Payment
curl http://localhost:8085/health  # Management
```

Expected response for each:
```json
{"status":"healthy","service":"gateway"}
```

## Quick Test

### 1. Create a product
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Laptop", "price": 999.99}'
```

### 2. Add to cart
```bash
curl -X POST http://localhost:8080/cart/items \
  -H "Content-Type: application/json" \
  -H "X-User-Id: user-1" \
  -d '{"product_id": 1, "quantity": 1, "price": 999.99}'
```

### 3. Checkout (Full Flow)
```bash
curl -X POST http://localhost:8080/cart/checkout \
  -H "X-User-Id: user-1"
```

This triggers:
1. Gateway → Cart Service
2. Cart → Payment Service
3. Cart → Invoice Service
4. Invoice → Management Service

## Stopping Services

**Stop all services:**
```bash
pkill -f "go run main.go"
```

**Stop databases:**
```bash
docker-compose down
```

**Stop and remove all data:**
```bash
docker-compose down -v
```

## Troubleshooting

### Database Connection Issues

Check if PostgreSQL containers are running:
```bash
docker ps
```

Restart databases:
```bash
docker-compose restart
```

### Port Already in Use

Check what's using the port:
```bash
lsof -i :8080  # Replace with your port
```

Kill the process:
```bash
kill -9 <PID>
```

### Service Not Starting

Check logs:
```bash
# If running in background, check process status
ps aux | grep "go run"

# View Docker logs
docker-compose logs
```

## Next Steps

- See **API_EXAMPLES.md** for detailed API testing examples
- See **README.md** for architecture documentation
- Modify `docker-compose.yml` to change database ports
- Update service configs in each `config/config.go` file

## Architecture Reminder

```
Client
  ↓
Gateway (8080)
  ↓
  ├── Product Service (8081)
  ├── Cart Service (8082) → Payment Service (8084)
  │                       → Invoice Service (8083) → Management Service (8085)
  └── Invoice Service (8083)
```

## Fixed Users

- `user-1` → Alice
- `user-2` → Bob
- `user-3` → Charlie
- `user-4` → Diana

Use header: `X-User-Id: user-1`
