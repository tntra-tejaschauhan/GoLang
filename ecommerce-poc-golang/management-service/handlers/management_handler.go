package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"management-service/models"
	"management-service/repository"
)

type ManagementHandler struct {
	repo *repository.ManagementRepository
}

func NewManagementHandler(repo *repository.ManagementRepository) *ManagementHandler {
	return &ManagementHandler{repo: repo}
}

func (h *ManagementHandler) RegisterInvoice(c *gin.Context) {
	var req models.RegisterInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice := &models.RegisteredInvoice{
		InvoiceID: req.InvoiceID,
		UserID:    req.UserID,
		Amount:    req.Amount,
	}

	if err := h.repo.RegisterInvoice(invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register invoice"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Invoice registered successfully",
		"data":    invoice,
	})
}

func (h *ManagementHandler) ListRegisteredInvoices(c *gin.Context) {
	invoices, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoices"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": invoices,
	})
}

func (h *ManagementHandler) GetRegisteredInvoice(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	invoice, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": invoice,
	})
}
