package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"management-service/config"
	"management-service/handlers"
	"management-service/models"
	"management-service/repository"
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
	if err := db.AutoMigrate(&models.RegisteredInvoice{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repository and handler
	repo := repository.NewManagementRepository(db)
	handler := handlers.NewManagementHandler(repo)

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "management"})
	})

	// Management routes
	r.POST("/invoices/register", handler.RegisterInvoice)
	r.GET("/invoices/registered", handler.ListRegisteredInvoices)
	r.GET("/invoices/registered/:id", handler.GetRegisteredInvoice)

	log.Printf("Management service starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start management service: %v", err)
	}
}
