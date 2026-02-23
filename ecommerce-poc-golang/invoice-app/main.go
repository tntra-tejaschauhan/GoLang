package start

import (
	"log"

	"invoice-app/config"
	"invoice-app/handlers"
	"invoice-app/models"
	"invoice-app/repository"
	"invoice-app/services"

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
	if err := db.AutoMigrate(&models.Invoice{}, &models.InvoiceItem{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize layers
	invoiceRepo := repository.NewInvoiceRepository(db)
	managementService := services.NewManagementService(cfg)
	handler := handlers.NewInvoiceHandler(invoiceRepo, managementService)

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "invoice"})
	})

	// Invoice routes
	r.POST("/invoices", handler.CreateInvoice)
	r.GET("/invoices/:id", handler.GetInvoice)
	r.GET("/invoices", handler.ListInvoices)

	log.Printf("Invoice service starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start invoice service: %v", err)
	}
}
