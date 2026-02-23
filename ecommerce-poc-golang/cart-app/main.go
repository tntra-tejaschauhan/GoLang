package main

import (
	"log"

	"cart-app/config"
	"cart-app/handlers"
	"cart-app/models"
	"cart-app/repository"
	"cart-app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	if err := db.AutoMigrate(&models.Cart{}, &models.CartItem{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize layers
	cartRepo := repository.NewCartRepository(db)
	checkoutService := services.NewCheckoutService(cfg, cartRepo)
	handler := handlers.NewCartHandler(cartRepo, checkoutService)

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "cart"})
	})

	// Cart routes
	r.POST("/cart/items", handler.AddItem)
	r.GET("/cart", handler.GetCart)
	r.POST("/cart/checkout", handler.Checkout)

	log.Printf("Cart service starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start cart service: %v", err)
	}
}
