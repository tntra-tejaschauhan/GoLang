package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	serviceURL string
	client     *http.Client
}

func NewProductHandler(serviceURL string) *ProductHandler {
	return &ProductHandler{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	req, _ := http.NewRequest("POST", h.serviceURL+"/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

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

func (h *ProductHandler) ListProducts(c *gin.Context) {
	req, _ := http.NewRequest("GET", h.serviceURL+"/products", nil)

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

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	req, _ := http.NewRequest("GET", h.serviceURL+"/products/"+id, nil)

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
