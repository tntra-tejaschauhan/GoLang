# Modular Monolith with Internal Gateway Architecture

## Overview

This project follows a **Modular Monolith architecture** using a **Single Entry Gateway pattern**.

All services run inside **one application**, exposed through **one port**, and all external requests must pass through a centralized **Gateway layer**.

---

## Architecture Flow

Client
   ↓
Gateway (Routing Layer)
   ↓
Controller
   ↓
Service
   ↓
Repository
   ↓
Database

---

## Core Concepts

### 1. Routing

Routing is the process of mapping an incoming HTTP request (method + URL) to the correct controller/handler.

Routing matches:
- HTTP Method (GET, POST, PUT, DELETE)
- URL Path (/users, /orders)
- Path parameters (/users/{id})

Routing does NOT contain business logic.

---

### 2. Gateway Pattern (Single Entry Point)

All external requests must go through the Gateway.

Services:
- Do NOT expose their own ports
- Do NOT start their own server
- Do NOT self-register routes

Gateway responsibilities:
- Register all public routes
- Apply middleware (authentication, logging, validation)
- Delegate to controllers and services

---

## Layer Responsibilities

### Gateway Layer
- Defines URL paths
- Maps URL → Controller
- Applies middleware
- Acts as entry point

### Controller Layer
- Handles HTTP request/response
- Validates input
- Calls service layer
- Contains no business logic

### Service Layer
- Contains business logic
- Calls repository
- Independent of HTTP

### Repository Layer
- Handles database operations
- Pure data access logic

---

## Important Architectural Rules

### Inside a Monolith

- Use direct function calls between modules
- No internal HTTP calls
- No URL forwarding
- No controller-to-controller communication

### Avoid

- Calling internal services via HTTP
- Letting modules expose routes independently
- Mixing business logic inside controllers

---

## Spring Boot Structure (Modular Monolith)

Project Structure:

myapp/
│
├── gateway-app
├── user-module
├── order-module
├── payment-module
└── inventory-module

Key Points:
- Multi-module Maven/Gradle project
- One Spring Boot application
- One DispatcherServlet (router)
- Gateway registers public controllers
- Modules contain services and repositories

Spring Flow:

Client
  ↓
DispatcherServlet
  ↓
GatewayController
  ↓
UserService (module)

---

## Go Structure (Modular Monolith)

Project Structure:

myapp/
│
├── cmd/main.go
│
├── internal/
│   ├── gateway/
│   ├── user/
│   ├── order/  
│   ├── payment/
│   └── inventory/
│
└── go.mod

Key Points:
- Single Go module
- Multiple internal packages
- One binary
- One running process
- Gateway registers all routes
- Services do not expose themselves

Go Flow:

Client
  ↓
Router (ServeMux / Gin)
  ↓
Controller
  ↓
Service

---

## Modular Monolith vs Microservices

| Modular Monolith | Microservices |
|------------------|--------------|
| One process | Multiple processes |
| One port | Multiple ports |
| Internal function calls | HTTP/gRPC between services |
| Easier development | More complex architecture |
| Suitable for most systems | Suitable for very large scale |

---

## Why Gateway-Only Access?

- Centralized security
- Centralized logging
- Centralized rate limiting
- Consistent API structure
- Easier migration to microservices in the future

---

## Final Architecture Statement

This system follows a **Modular Monolith architecture with a Single Entry Gateway pattern**.

All external requests are routed through a centralized Gateway layer, which maps URL paths to controllers. Controllers delegate business logic to service modules via direct function calls. No internal HTTP communication exists between modules. The application runs as a single process with clearly separated layers to ensure scalability, maintainability, and future microservice readiness.




```
ecommerce-parent/
│
├── pom.xml  (Parent POM)
│
├── gateway-app/
│   ├── pom.xml
│   └── src/main/java/com/example/gateway/
│       ├── GatewayApplication.java
│       ├── controller/
│       │     ├── ProductGatewayController.java
│       │     ├── CartGatewayController.java
│       │     └── InvoiceGatewayController.java
│       └── config/
│             └── WebConfig.java
│
├── product-module/
│   └── src/main/java/com/example/product/
│       ├── service/
│       │     └── ProductService.java
│       ├── repository/
│       │     └── ProductRepository.java
│       └── model/
│             └── Product.java
│
├── cart-module/
│   └── src/main/java/com/example/cart/
│       ├── service/
│       │     └── CartService.java
│       ├── repository/
│       │     └── CartRepository.java
│       └── model/
│             └── Cart.java
│
└── invoice-module/
    └── src/main/java/com/example/invoice/
        ├── service/
        │     └── InvoiceService.java
        ├── repository/
        │     └── InvoiceRepository.java
        └── model/
              └── Invoice.java

```

