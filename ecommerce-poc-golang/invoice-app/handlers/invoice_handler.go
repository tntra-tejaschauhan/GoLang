package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"invoice-app/models"
	"invoice-app/repository"
	"invoice-app/services"
)

type InvoiceHandler struct {
	invoiceRepo       *repository.InvoiceRepository
	managementService *services.ManagementService
}

func NewInvoiceHandler(invoiceRepo *repository.InvoiceRepository, managementService *services.ManagementService) *InvoiceHandler {
	return &InvoiceHandler{
		invoiceRepo:       invoiceRepo,
		managementService: managementService,
	}
}

func (h *InvoiceHandler) CreateInvoice(c *gin.Context) {
	var req models.CreateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create invoice
	invoice := &models.Invoice{
		UserID:      req.UserID,
		TotalAmount: req.TotalAmount,
		PaymentID:   req.PaymentID,
	}

	// Create invoice items
	for _, item := range req.Items {
		invoice.Items = append(invoice.Items, models.InvoiceItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	if err := h.invoiceRepo.Create(invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		return
	}

	// Register with Management Service
	if err := h.managementService.RegisterInvoice(invoice.ID, invoice.UserID, invoice.TotalAmount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register invoice with management"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Invoice created successfully",
		"data":    invoice,
	})
}

func (h *InvoiceHandler) GetInvoice(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Id header is required"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	invoice, err := h.invoiceRepo.FindByID(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": invoice,
	})
}

func (h *InvoiceHandler) ListInvoices(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Id header is required"})
		return
	}

	invoices, err := h.invoiceRepo.FindByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoices"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": invoices,
	})
}
