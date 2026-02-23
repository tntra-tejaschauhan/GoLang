package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gateway-app/config"
	"gateway-app/handlers"
	"gateway-app/middleware"
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	// Middleware
	r.Use(middleware.UserValidator())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "gateway"})
	})

	// Initialize handlers
	productHandler := handlers.NewProductHandler(cfg.ProductServiceURL)
	cartHandler := handlers.NewCartHandler(cfg.CartServiceURL)
	invoiceHandler := handlers.NewInvoiceHandler(cfg.InvoiceServiceURL)

	// Product routes
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.ListProducts)
	r.GET("/products/:id", productHandler.GetProduct)

	// Cart routes (requires X-User-Id)
	r.POST("/cart/items", cartHandler.AddItem)
	r.GET("/cart", cartHandler.GetCart)
	r.POST("/cart/checkout", cartHandler.Checkout)

	// Invoice routes (requires X-User-Id)
	r.GET("/invoices/:id", invoiceHandler.GetInvoice)
	r.GET("/invoices", invoiceHandler.ListInvoices)

	log.Printf("Gateway starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start gateway: %v", err)
	}
}
