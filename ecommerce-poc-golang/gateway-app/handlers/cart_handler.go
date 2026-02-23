package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	serviceURL string
	client     *http.Client
}

func NewCartHandler(serviceURL string) *CartHandler {
	return &CartHandler{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

func (h *CartHandler) AddItem(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")
	
	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	req, _ := http.NewRequest("POST", h.serviceURL+"/cart/items", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User-Id", userID)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	c.JSON(resp.StatusCode, result)
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")

	req, _ := http.NewRequest("GET", h.serviceURL+"/cart", nil)
	req.Header.Set("X-User-Id", userID)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	c.JSON(resp.StatusCode, result)
}

func (h *CartHandler) Checkout(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")

	req, _ := http.NewRequest("POST", h.serviceURL+"/cart/checkout", nil)
	req.Header.Set("X-User-Id", userID)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	c.JSON(resp.StatusCode, result)
}
