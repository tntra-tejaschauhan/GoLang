package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"payment-service/config"
	"payment-service/models"
	"payment-service/repository"
)

type PaymentHandler struct {
	repo   *repository.PaymentRepository
	config *config.Config
}

func NewPaymentHandler(repo *repository.PaymentRepository, cfg *config.Config) *PaymentHandler {
	return &PaymentHandler{
		repo:   repo,
		config: cfg,
	}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req models.ProcessPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Business rule: if amount < limit → SUCCESS, else → FAILED
	status := "SUCCESS"
	if req.Amount >= h.config.PaymentLimit {
		status = "FAILED"
	}

	payment := &models.Payment{
		Amount: req.Amount,
		Status: status,
	}

	if err := h.repo.Create(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	statusCode := http.StatusOK
	if status == "FAILED" {
		statusCode = http.StatusPaymentRequired
	}

	c.JSON(statusCode, gin.H{
		"message": "Payment processed",
		"data":    payment,
	})
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	payment, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payment,
	})
}
