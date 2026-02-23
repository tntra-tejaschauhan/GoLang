package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"product-app/config"
	"product-app/handlers"
	"product-app/models"
	"product-app/repository"
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
	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repository and handler
	repo := repository.NewProductRepository(db)
	handler := handlers.NewProductHandler(repo)

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "product"})
	})

	// Product routes
	r.POST("/products", handler.CreateProduct)
	r.GET("/products", handler.ListProducts)
	r.GET("/products/:id", handler.GetProduct)

	log.Printf("Product service starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start product service: %v", err)
	}
}
