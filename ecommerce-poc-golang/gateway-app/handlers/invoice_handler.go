package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	serviceURL string
	client     *http.Client
}

func NewInvoiceHandler(serviceURL string) *InvoiceHandler {
	return &InvoiceHandler{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

func (h *InvoiceHandler) GetInvoice(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")
	id := c.Param("id")

	req, _ := http.NewRequest("GET", h.serviceURL+"/invoices/"+id, nil)
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

func (h *InvoiceHandler) ListInvoices(c *gin.Context) {
	userID := c.GetHeader("X-User-Id")

	req, _ := http.NewRequest("GET", h.serviceURL+"/invoices", nil)
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
