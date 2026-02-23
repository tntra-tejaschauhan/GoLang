# Architecture Documentation

## System Overview

This is a modular e-commerce POC system built with Go, following microservices-inspired patterns while maintaining simplicity for POC purposes.

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────────┐
│                              CLIENT                                  │
└────────────────────────────────┬────────────────────────────────────┘
                                 │
                                 ▼
                    ┌────────────────────────┐
                    │   GATEWAY SERVICE      │
                    │   (Port 8080)          │
                    │   - User Validation    │
                    │   - Request Routing    │
                    └────────────┬───────────┘
                                 │
                ┌────────────────┼────────────────┐
                │                │                │
                ▼                ▼                ▼
    ┌──────────────────┐ ┌──────────────┐ ┌────────────────┐
    │ PRODUCT SERVICE  │ │ CART SERVICE │ │ INVOICE SERVICE│
    │  (Port 8081)     │ │ (Port 8082)  │ │  (Port 8083)   │
    │                  │ │              │ │                │
    │ - Product CRUD   │ │ - Add Items  │ │ - Create Invoice│
    │ - Product List   │ │ - View Cart  │ │ - View Invoice │
    └────────┬─────────┘ │ - Checkout   │ └────────┬───────┘
             │           └──────┬───────┘          │
             │                  │                  │
             ▼                  │                  │
    ┌──────────────────┐       │                  │
    │  PostgreSQL      │       │                  │
    │  product_db      │       │                  │
    │  (Port 5432)     │       │                  │
    └──────────────────┘       │                  │
                                │                  │
                                ▼                  ▼
                    ┌────────────────────┐  ┌──────────────────┐
                    │ PAYMENT SERVICE    │  │ PostgreSQL       │
                    │ (Port 8084)        │  │ invoice_db       │
                    │ [EXTERNAL]         │  │ (Port 5434)      │
                    │                    │  └──────────────────┘
                    │ - Process Payment  │
                    │ - Payment Limit    │
                    │   Check ($1000)    │
                    └──────────┬─────────┘
                               │
                               ▼
                    ┌──────────────────┐
                    │  PostgreSQL      │
                    │  payment_db      │
                    │  (Port 5435)     │
                    └──────────────────┘
                               
                               
    ┌──────────────────┐
    │  PostgreSQL      │◄──────────┘
    │  cart_db         │
    │  (Port 5433)     │
    └──────────────────┘


                    ┌────────────────────────┐
                    │  MANAGEMENT SERVICE    │
                    │  (Port 8085)           │
                    │  [EXTERNAL]            │
                    │                        │
                    │  - Register Invoice    │
                    │  - ACK                 │
                    └───────────┬────────────┘
                                │
                                ▼
                    ┌──────────────────────┐
                    │  PostgreSQL          │
                    │  management_db       │
                    │  (Port 5436)         │
                    └──────────────────────┘
```

## Request Flow

### Complete Checkout Flow

```
1. Client → Gateway
   POST /cart/checkout
   Header: X-User-Id: user-1

2. Gateway → Cart Service
   POST /cart/checkout
   Header: X-User-Id: user-1

3. Cart Service → Payment Service (HTTP Call)
   POST /payments
   Body: { "amount": 1999.98 }
   
   Response: { "id": 1, "status": "SUCCESS" }

4. Cart Service → Invoice Service (HTTP Call)
   POST /invoices
   Body: {
     "user_id": "user-1",
     "total_amount": 1999.98,
     "payment_id": 1,
     "items": [...]
   }
   
   Response: { "id": 1, ... }

5. Invoice Service → Management Service (HTTP Call)
   POST /invoices/register
   Body: {
     "invoice_id": 1,
     "user_id": "user-1",
     "amount": 1999.98
   }
   
   Response: { "status": "registered" }

6. Cart Service → Gateway → Client
   Response: {
     "message": "Checkout completed",
     "cart_id": 1,
     "payment_id": 1,
     "invoice_id": 1,
     "total": 1999.98
   }
```

## Service Boundaries

### Internal Services (Modular Monolith Pattern)
- **Gateway Service**: Entry point, user validation, routing
- **Product Service**: Product catalog management
- **Cart Service**: Shopping cart and checkout orchestration
- **Invoice Service**: Invoice creation and management

### External Services (Separate Applications)
- **Payment Service**: Payment processing (simulated)
- **Management Service**: Invoice registration and tracking

## Database Strategy

Each service has its own PostgreSQL database:

| Service    | Database         | Port | Tables                    |
|------------|------------------|------|---------------------------|
| Product    | product_db       | 5432 | products                  |
| Cart       | cart_db          | 5433 | carts, cart_items         |
| Invoice    | invoice_db       | 5434 | invoices, invoice_items   |
| Payment    | payment_db       | 5435 | payments                  |
| Management | management_db    | 5436 | registered_invoices       |

## Communication Patterns

### Synchronous HTTP Communication

All inter-service communication uses HTTP REST APIs:
- Gateway → Internal Services: HTTP
- Cart → Payment: HTTP
- Cart → Invoice: HTTP
- Invoice → Management: HTTP

### Headers

- **X-User-Id**: Required for cart and invoice operations
- **Content-Type**: application/json for all POST requests

## Data Models

### Product Service
```go
type Product struct {
    ID        uint
    Name      string
    Price     float64
    CreatedAt time.Time
}
```

### Cart Service
```go
type Cart struct {
    ID     uint
    UserID string
    Status string  // "active", "checked_out"
    Items  []CartItem
}

type CartItem struct {
    ID        uint
    CartID    uint
    ProductID uint
    Quantity  int
    Price     float64
}
```

### Invoice Service
```go
type Invoice struct {
    ID          uint
    UserID      string
    TotalAmount float64
    PaymentID   uint
    Items       []InvoiceItem
}

type InvoiceItem struct {
    ID        uint
    InvoiceID uint
    ProductID uint
    Quantity  int
    Price     float64
}
```

### Payment Service
```go
type Payment struct {
    ID     uint
    Amount float64
    Status string  // "SUCCESS", "FAILED"
}
```

### Management Service
```go
type RegisteredInvoice struct {
    ID         uint
    InvoiceID  uint
    UserID     string
    Amount     float64
    ReceivedAt time.Time
}
```

## Business Rules

### Payment Service
- If amount < $1000 → Payment SUCCESS
- If amount >= $1000 → Payment FAILED

### Cart Service
- One active cart per user
- Cart status changes to "checked_out" after successful checkout
- Checkout fails if payment is not successful

### User Management
- Fixed 4 users (no authentication):
  - user-1 → Alice
  - user-2 → Bob
  - user-3 → Charlie
  - user-4 → Diana

## Technology Stack

| Component        | Technology              |
|------------------|-------------------------|
| Language         | Go 1.21+                |
| Web Framework    | Gin                     |
| ORM             | GORM                     |
| Database         | PostgreSQL 15           |
| Containerization | Docker, Docker Compose  |
| API Style        | REST                    |

## Error Handling

Each service returns standardized JSON responses:

**Success:**
```json
{
  "message": "Operation successful",
  "data": { ... }
}
```

**Error:**
```json
{
  "error": "Error description"
}
```

## Scalability Considerations

### Current POC
- All services can run independently
- Separate databases per service
- HTTP-based communication

### Future Enhancements
- Add API Gateway (Kong, Traefik)
- Implement circuit breakers
- Add distributed tracing (Jaeger)
- Implement event-driven architecture (Kafka, RabbitMQ)
- Add caching layer (Redis)
- Implement authentication (JWT)
- Add service mesh (Istio)

## Development Best Practices

1. **Separation of Concerns**: Each service has clear boundaries
2. **Repository Pattern**: Data access abstraction
3. **Handler-Service-Repository**: Three-layer architecture
4. **Configuration Management**: Environment-based configs
5. **Database Migrations**: Automatic with GORM AutoMigrate
6. **Health Checks**: Each service exposes /health endpoint

## Testing Strategy

### Unit Tests
Test individual functions in handlers, services, and repositories

### Integration Tests
Test HTTP endpoints with actual database

### End-to-End Tests
Test complete checkout flow through gateway

## Monitoring & Observability

Consider adding:
- Prometheus for metrics
- Grafana for dashboards
- ELK stack for logging
- Distributed tracing

## Security Considerations (For Production)

⚠️ This POC does not implement:
- Authentication/Authorization
- TLS/HTTPS
- Rate limiting
- Input validation (minimal)
- SQL injection prevention (GORM handles this)
- CORS configuration

## Deployment Options

### Development
- Run locally with `go run`
- Use Docker Compose for databases

### Production (Future)
- Containerize each service
- Deploy to Kubernetes
- Use managed PostgreSQL (AWS RDS, Azure PostgreSQL)
- Implement CI/CD pipeline
- Add load balancers

## Performance Optimization

For production:
- Add connection pooling
- Implement caching strategies
- Use prepared statements
- Add database indexes
- Implement pagination
- Use goroutines for parallel HTTP calls

## Maintenance

### Adding New Features
1. Identify which service owns the feature
2. Create models, repositories, handlers
3. Update routes
4. Test independently
5. Update documentation

### Database Changes
1. Modify model structs
2. GORM will auto-migrate on startup
3. For production, use proper migrations

## Conclusion

This architecture balances:
- **Simplicity**: Easy to understand and develop
- **Modularity**: Clear service boundaries
- **Scalability**: Can evolve into microservices
- **Maintainability**: Organized codebase
- **Testability**: Services can be tested independently
