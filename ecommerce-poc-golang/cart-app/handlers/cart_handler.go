package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"cart-app/models"
	"cart-app/repository"
	"cart-app/services"
)

type CartHandler struct {
	cartRepo        *repository.CartRepository
	checkoutService *services.CheckoutService
}

func NewCartHandler(cartRepo *repository.CartRepository, checkoutService *services.CheckoutService) *CartHandler {
	return &CartHandler{
		cartRepo:        cartRepo,
		checkoutService: checkoutService,
	}
}

func (h *CartHandler) AddItem(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Id header is required"})
		return
	}

	var req models.AddItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get or create active cart
	cart, err := h.cartRepo.GetOrCreateActiveCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart"})
		return
	}

	// Add item
	item := &models.CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Price:     req.Price,
	}

	if err := h.cartRepo.AddItem(cart.ID, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Item added to cart",
		"data":    item,
	})
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Id header is required"})
		return
	}

	cart, err := h.cartRepo.GetOrCreateActiveCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart"})
		return
	}

	total, _ := h.cartRepo.CalculateTotal(cart.ID)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"cart":  cart,
			"total": total,
		},
	})
}

func (h *CartHandler) Checkout(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Id header is required"})
		return
	}

	result, err := h.checkoutService.ProcessCheckout(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
