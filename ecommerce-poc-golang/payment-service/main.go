package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"payment-service/config"
	"payment-service/handlers"
	"payment-service/models"
	"payment-service/repository"
)

func main() {
	cfg := config.Load()

	// Database connection
	dsn := cfg.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(&models.Payment{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repository and handler
	repo := repository.NewPaymentRepository(db)
	handler := handlers.NewPaymentHandler(repo, cfg)

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "payment"})
	})

	// Payment routes
	r.POST("/payments", handler.ProcessPayment)
	r.GET("/payments/:id", handler.GetPayment)

	log.Printf("Payment service starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start payment service: %v", err)
	}
}
