# API Testing Examples

## Setup
Make sure all services are running and databases are up.

## 1. Create a Product

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "price": 999.99
  }'
```

Response:
```json
{
  "message": "Product created successfully",
  "data": {
    "id": 1,
    "name": "Laptop",
    "price": 999.99,
    "created_at": "2024-01-15T10:00:00Z"
  }
}
```

## 2. List Products

```bash
curl http://localhost:8080/products
```

## 3. Add Item to Cart

```bash
curl -X POST http://localhost:8080/cart/items \
  -H "Content-Type: application/json" \
  -H "X-User-Id: user-1" \
  -d '{
    "product_id": 1,
    "quantity": 2,
    "price": 999.99
  }'
```

## 4. View Cart

```bash
curl http://localhost:8080/cart \
  -H "X-User-Id: user-1"
```

Response:
```json
{
  "data": {
    "cart": {
      "id": 1,
      "user_id": "user-1",
      "status": "active",
      "items": [
        {
          "id": 1,
          "product_id": 1,
          "quantity": 2,
          "price": 999.99
        }
      ]
    },
    "total": 1999.98
  }
}
```

## 5. Checkout (End-to-End Flow)

```bash
curl -X POST http://localhost:8080/cart/checkout \
  -H "X-User-Id: user-1"
```

This triggers the full flow:
- Gateway → Cart Service
- Cart Service → Payment Service (HTTP)
- Cart Service → Invoice Service (HTTP)
- Invoice Service → Management Service (HTTP)

Response:
```json
{
  "data": {
    "message": "Checkout completed successfully",
    "cart_id": 1,
    "total": 1999.98,
    "payment_id": 1,
    "invoice_id": 1
  }
}
```

## 6. View Invoice

```bash
curl http://localhost:8080/invoices/1 \
  -H "X-User-Id: user-1"
```

## 7. Test Payment Limit

Create a high-value product:
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Luxury Car",
    "price": 50000.00
  }'
```

Add to cart and checkout - payment will FAIL (amount >= 1000):
```bash
curl -X POST http://localhost:8080/cart/items \
  -H "Content-Type: application/json" \
  -H "X-User-Id: user-2" \
  -d '{
    "product_id": 2,
    "quantity": 1,
    "price": 50000.00
  }'

curl -X POST http://localhost:8080/cart/checkout \
  -H "X-User-Id: user-2"
```

## Direct Service Access (for testing)

### Product Service (Direct)
```bash
curl http://localhost:8081/products
```

### Cart Service (Direct)
```bash
curl http://localhost:8082/cart \
  -H "X-User-Id: user-1"
```

### Payment Service (Direct)
```bash
curl -X POST http://localhost:8084/payments \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 500.00
  }'
```

### Management Service (Direct)
```bash
curl http://localhost:8085/invoices/registered
```

## Health Checks

```bash
curl http://localhost:8080/health  # Gateway
curl http://localhost:8081/health  # Product
curl http://localhost:8082/health  # Cart
curl http://localhost:8083/health  # Invoice
curl http://localhost:8084/health  # Payment
curl http://localhost:8085/health  # Management
```

## Complete Test Scenario

```bash
# 1. Create products
curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d '{"name": "Mouse", "price": 25.99}'
curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d '{"name": "Keyboard", "price": 75.50}'

# 2. User-1: Add items to cart
curl -X POST http://localhost:8080/cart/items -H "Content-Type: application/json" -H "X-User-Id: user-1" -d '{"product_id": 1, "quantity": 1, "price": 25.99}'
curl -X POST http://localhost:8080/cart/items -H "Content-Type: application/json" -H "X-User-Id: user-1" -d '{"product_id": 2, "quantity": 2, "price": 75.50}'

# 3. View cart
curl http://localhost:8080/cart -H "X-User-Id: user-1"

# 4. Checkout (triggers payment and invoice)
curl -X POST http://localhost:8080/cart/checkout -H "X-User-Id: user-1"

# 5. View invoice
curl http://localhost:8080/invoices/1 -H "X-User-Id: user-1"

# 6. View all registered invoices in management
curl http://localhost:8085/invoices/registered
```
