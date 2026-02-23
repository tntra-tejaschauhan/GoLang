package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"invoice-app/config"
)

type ManagementService struct {
	config *config.Config
	client *http.Client
}

func NewManagementService(cfg *config.Config) *ManagementService {
	return &ManagementService{
		config: cfg,
		client: &http.Client{},
	}
}

type RegisterInvoiceRequest struct {
	InvoiceID uint    `json:"invoice_id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
}

type RegisterInvoiceResponse struct {
	ID        uint    `json:"id"`
	InvoiceID uint    `json:"invoice_id"`
	Status    string  `json:"status"`
}

func (s *ManagementService) RegisterInvoice(invoiceID uint, userID string, amount float64) error {
	req := RegisterInvoiceRequest{
		InvoiceID: invoiceID,
		UserID:    userID,
		Amount:    amount,
	}

	jsonData, _ := json.Marshal(req)

	httpReq, _ := http.NewRequest("POST", s.config.ManagementServiceURL+"/invoices/register", bytes.NewBuffer(jsonData))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to register invoice: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return fmt.Errorf("management service returned status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
