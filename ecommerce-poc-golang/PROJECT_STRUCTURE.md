# Project Structure

```
ecommerce-poc-golang/
│
├── README.md                    # Main documentation
├── QUICKSTART.md                # Quick start guide
├── ARCHITECTURE.md              # Architecture documentation
├── API_EXAMPLES.md              # API testing examples
├── docker-compose.yml           # PostgreSQL databases setup
├── setup.sh                     # Setup script
│
├── shared/                      # Shared utilities (optional)
│   ├── database/
│   │   └── connection.go       # Database connection utilities
│   ├── httpclient/
│   │   └── client.go           # HTTP client utilities
│   └── models/
│       └── common.go           # Common data structures
│
├── gateway-app/                 # API Gateway Service
│   ├── main.go                 # Entry point
│   ├── go.mod                  # Go dependencies
│   ├── config/
│   │   └── config.go          # Configuration
│   ├── middleware/
│   │   └── user_validator.go # User validation middleware
│   └── handlers/
│       ├── product_handler.go # Product proxy
│       ├── cart_handler.go    # Cart proxy
│       └── invoice_handler.go # Invoice proxy
│
├── product-app/                 # Product Service
│   ├── main.go                 # Entry point
│   ├── go.mod                  # Go dependencies
│   ├── config/
│   │   └── config.go          # Configuration
│   ├── models/
│   │   └── product.go         # Product model
│   ├── repository/
│   │   └── product_repository.go  # Data access
│   └── handlers/
│       └── product_handler.go     # HTTP handlers
│
├── cart-app/                    # Cart Service
│   ├── main.go                 # Entry point
│   ├── go.mod                  # Go dependencies
│   ├── config/
│   │   └── config.go          # Configuration
│   ├── models/
│   │   └── cart.go            # Cart & CartItem models
│   ├── repository/
│   │   └── cart_repository.go # Data access
│   ├── services/
│   │   └── checkout_service.go # Checkout orchestration
│   └── handlers/
│       └── cart_handler.go    # HTTP handlers
│
├── invoice-app/                 # Invoice Service
│   ├── main.go                 # Entry point
│   ├── go.mod                  # Go dependencies
│   ├── config/
│   │   └── config.go          # Configuration
│   ├── models/
│   │   └── invoice.go         # Invoice & InvoiceItem models
│   ├── repository/
│   │   └── invoice_repository.go  # Data access
│   ├── services/
│   │   └── management_service.go  # Management client
│   └── handlers/
│       └── invoice_handler.go     # HTTP handlers
│
├── payment-service/             # Payment Service (External)
│   ├── main.go                 # Entry point
│   ├── go.mod                  # Go dependencies
│   ├── config/
│   │   └── config.go          # Configuration
│   ├── models/
│   │   └── payment.go         # Payment model
│   ├── repository/
│   │   └── payment_repository.go  # Data access
│   └── handlers/
│       └── payment_handler.go     # HTTP handlers
│
└── management-service/          # Management Service (External)
    ├── main.go                 # Entry point
    ├── go.mod                  # Go dependencies
    ├── config/
    │   └── config.go          # Configuration
    ├── models/
    │   └── registered_invoice.go  # RegisteredInvoice model
    ├── repository/
    │   └── management_repository.go  # Data access
    └── handlers/
        └── management_handler.go     # HTTP handlers
```

## Service Details

### Gateway Service (Port 8080)
**Purpose**: API Gateway and request router  
**Dependencies**: None (stateless)  
**Database**: None  
**Key Files**:
- `middleware/user_validator.go`: Validates X-User-Id header
- `handlers/*_handler.go`: Proxy requests to backend services

### Product Service (Port 8081)
**Purpose**: Product catalog management  
**Dependencies**: None  
**Database**: `product_db` (port 5432)  
**Key Files**:
- `models/product.go`: Product entity
- `repository/product_repository.go`: CRUD operations
- `handlers/product_handler.go`: REST endpoints

**Endpoints**:
- `POST /products` - Create product
- `GET /products` - List products
- `GET /products/:id` - Get product

### Cart Service (Port 8082)
**Purpose**: Shopping cart and checkout orchestration  
**Dependencies**: 
- Payment Service (HTTP)
- Invoice Service (HTTP)  
**Database**: `cart_db` (port 5433)  
**Key Files**:
- `models/cart.go`: Cart and CartItem entities
- `services/checkout_service.go`: Orchestrates payment and invoice
- `handlers/cart_handler.go`: REST endpoints

**Endpoints**:
- `POST /cart/items` - Add item
- `GET /cart` - View cart
- `POST /cart/checkout` - Checkout

### Invoice Service (Port 8083)
**Purpose**: Invoice creation and management  
**Dependencies**: Management Service (HTTP)  
**Database**: `invoice_db` (port 5434)  
**Key Files**:
- `models/invoice.go`: Invoice and InvoiceItem entities
- `services/management_service.go`: Management client
- `handlers/invoice_handler.go`: REST endpoints

**Endpoints**:
- `POST /invoices` - Create invoice
- `GET /invoices/:id` - Get invoice
- `GET /invoices` - List user invoices

### Payment Service (Port 8084) - EXTERNAL
**Purpose**: Payment processing  
**Dependencies**: None  
**Database**: `payment_db` (port 5435)  
**Business Rule**: Amount < $1000 → SUCCESS, else FAILED  
**Key Files**:
- `models/payment.go`: Payment entity
- `handlers/payment_handler.go`: REST endpoints

**Endpoints**:
- `POST /payments` - Process payment
- `GET /payments/:id` - Get payment

### Management Service (Port 8085) - EXTERNAL
**Purpose**: Invoice registration and tracking  
**Dependencies**: None  
**Database**: `management_db` (port 5436)  
**Key Files**:
- `models/registered_invoice.go`: RegisteredInvoice entity
- `handlers/management_handler.go`: REST endpoints

**Endpoints**:
- `POST /invoices/register` - Register invoice
- `GET /invoices/registered` - List all
- `GET /invoices/registered/:id` - Get by ID

## Configuration Files

### go.mod
Each service has its own `go.mod` with dependencies:
- `github.com/gin-gonic/gin` - Web framework
- `gorm.io/gorm` - ORM
- `gorm.io/driver/postgres` - PostgreSQL driver

### config/config.go
Each service has configuration for:
- Server port
- Database connection (host, port, user, password, database name)
- External service URLs (for cart and invoice services)

## Database Schema

### product_db
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### cart_db
```sql
CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INTEGER REFERENCES carts(id),
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### invoice_db
```sql
CREATE TABLE invoices (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    payment_id INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE invoice_items (
    id SERIAL PRIMARY KEY,
    invoice_id INTEGER REFERENCES invoices(id),
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### payment_db
```sql
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### management_db
```sql
CREATE TABLE registered_invoices (
    id SERIAL PRIMARY KEY,
    invoice_id INTEGER NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    received_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

## Architecture Layers

### Layer 1: Entry Point
- `main.go` - Application bootstrap, dependency injection

### Layer 2: HTTP Layer
- `handlers/` - HTTP request/response handling, validation

### Layer 3: Business Logic
- `services/` - Business logic, orchestration (only in cart and invoice)

### Layer 4: Data Access
- `repository/` - Database operations, queries

### Layer 5: Models
- `models/` - Data structures, entities

### Layer 6: Configuration
- `config/` - Application configuration, environment variables

## Design Patterns Used

1. **Repository Pattern**: Abstracts data access
2. **Dependency Injection**: Services receive dependencies via constructors
3. **Layered Architecture**: Clear separation of concerns
4. **Configuration Pattern**: Centralized config management
5. **Gateway Pattern**: Single entry point for clients

## File Naming Conventions

- `*_handler.go` - HTTP handlers
- `*_service.go` - Business logic services
- `*_repository.go` - Data access layer
- `*.go` - Models and entities
- `config.go` - Configuration

## Code Organization Best Practices

1. **One responsibility per file**
2. **Clear package structure**
3. **Consistent naming**
4. **Minimal dependencies**
5. **Easy to test**

## Development Workflow

1. **Add new endpoint**:
   - Add model in `models/`
   - Add repository method in `repository/`
   - Add handler in `handlers/`
   - Register route in `main.go`

2. **Add external service call**:
   - Create service in `services/`
   - Inject HTTP client
   - Use in handler

3. **Modify database**:
   - Update model struct
   - GORM auto-migrates on startup
