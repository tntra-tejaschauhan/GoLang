# E-Commerce POC - Golang Architecture

## 🎯 Goal

Build a POC e-commerce-like system using:
- ✅ Go 1.21+
- ✅ Gin Framework (HTTP server)
- ✅ PostgreSQL (per service DB)
- ✅ Go Modules for dependency management
- ✅ 4 fixed users (no authentication)
- ✅ Real HTTP communication between services
- ✅ Clear module boundaries

## 🧱 Architecture

```
Client
  |
  v
Modular System
  ├── gateway-app/          (Gin App)
  ├── product-app/          (Gin App)
  ├── cart-app/             (Gin App)  ───HTTP──▶ payment-service/ (Gin App)
  ├── invoice-app/          (Gin App) ───HTTP──▶ management-service/ (Gin App)
  └── shared/               (Common utilities)
```

## 📁 Project Structure

```
ecommerce-poc-golang/
├── gateway-app/
│   ├── main.go
│   ├── handlers/
│   ├── middleware/
│   ├── config/
│   └── go.mod
├── product-app/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   ├── migrations/
│   ├── config/
│   └── go.mod
├── cart-app/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   ├── services/
│   ├── migrations/
│   ├── config/
│   └── go.mod
├── invoice-app/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   ├── services/
│   ├── migrations/
│   ├── config/
│   └── go.mod
├── payment-service/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   ├── migrations/
│   ├── config/
│   └── go.mod
├── management-service/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   ├── migrations/
│   ├── config/
│   └── go.mod
├── shared/
│   ├── database/
│   ├── httpclient/
│   └── models/
├── docker-compose.yml
└── README.md
```

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL (via Docker)

### Setup

1. **Start PostgreSQL databases:**
```bash
docker-compose up -d
```

2. **Run each service** (in separate terminals):

```bash
# Gateway
cd gateway-app && go run main.go

# Product Service
cd product-app && go run main.go

# Cart Service
cd cart-app && go run main.go

# Invoice Service
cd invoice-app && go run main.go

# Payment Service (External)
cd payment-service && go run main.go

# Management Service (External)
cd management-service && go run main.go
```

## 🔌 Service Ports

- Gateway: `8080`
- Product: `8081`
- Cart: `8082`
- Invoice: `8083`
- Payment: `8084`
- Management: `8085`

## 👥 Fixed Users

```
user-1 → Alice
user-2 → Bob
user-3 → Charlie
user-4 → Diana
```

Send requests with header: `X-User-Id: user-1`

## 🔁 Request Flow

1. **Create Product** (Gateway)
   ```bash
   POST http://localhost:8080/products
   ```

2. **Add to Cart** (Gateway)
   ```bash
   POST http://localhost:8080/cart/items
   X-User-Id: user-1
   ```

3. **Checkout** (Gateway)
   ```bash
   POST http://localhost:8080/cart/checkout
   X-User-Id: user-1
   ```

   Flow:
   - Gateway → Cart App
   - Cart App → Payment Service (HTTP)
   - Cart App → Invoice App (HTTP)
   - Invoice App → Management Service (HTTP)

## 🗄️ Database Strategy

Each service has its own PostgreSQL database:
- `product_db`
- `cart_db`
- `invoice_db`
- `payment_db`
- `management_db`

## 📝 API Examples

See individual service READMEs for detailed API documentation.
