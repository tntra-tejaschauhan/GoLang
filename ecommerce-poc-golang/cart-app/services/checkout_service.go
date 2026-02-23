package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cart-app/config"
	"cart-app/models"
	"cart-app/repository"
)

type CheckoutService struct {
	config   *config.Config
	cartRepo *repository.CartRepository
	client   *http.Client
}

func NewCheckoutService(cfg *config.Config, cartRepo *repository.CartRepository) *CheckoutService {
	return &CheckoutService{
		config:   cfg,
		cartRepo: cartRepo,
		client:   &http.Client{},
	}
}

type PaymentRequest struct {
	Amount float64 `json:"amount"`
}

type PaymentResponse struct {
	ID     uint    `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

type InvoiceRequest struct {
	UserID    string               `json:"user_id"`
	Total     float64              `json:"total_amount"`
	PaymentID uint                 `json:"payment_id"`
	Items     []InvoiceItemRequest `json:"items"`
}

type InvoiceItemRequest struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type InvoiceResponse struct {
	ID        uint    `json:"id"`
	UserID    string  `json:"user_id"`
	Total     float64 `json:"total_amount"`
	PaymentID uint    `json:"payment_id"`
}

func (s *CheckoutService) ProcessCheckout(userID string) (*models.CheckoutResponse, error) {
	// 1. Get user's active cart
	cart, err := s.cartRepo.GetOrCreateActiveCart(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}

	if len(cart.Items) == 0 {
		return nil, fmt.Errorf("cart is empty")
	}

	// 2. Calculate total
	total, err := s.cartRepo.CalculateTotal(cart.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate total: %w", err)
	}

	// 3. Call Payment Service
	paymentResp, err := s.processPayment(total)
	if err != nil {
		return nil, fmt.Errorf("payment failed: %w", err)
	}

	if paymentResp.Status != "SUCCESS" {
		return nil, fmt.Errorf("payment declined")
	}

	// 4. Call Invoice Service
	invoiceResp, err := s.createInvoice(userID, total, paymentResp.ID, cart.Items)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}

	// 5. Mark cart as checked out
	if err := s.cartRepo.UpdateCartStatus(cart.ID, "checked_out"); err != nil {
		return nil, fmt.Errorf("failed to update cart status: %w", err)
	}

	return &models.CheckoutResponse{
		Message:   "Checkout completed successfully",
		CartID:    cart.ID,
		Total:     total,
		PaymentID: paymentResp.ID,
		InvoiceID: invoiceResp.ID,
	}, nil
}

func (s *CheckoutService) processPayment(amount float64) (*PaymentResponse, error) {
	paymentReq := PaymentRequest{Amount: amount}
	jsonData, _ := json.Marshal(paymentReq)

	req, _ := http.NewRequest("POST", s.config.PaymentServiceURL+"/payments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("payment service returned status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data PaymentResponse `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

func (s *CheckoutService) createInvoice(userID string, total float64, paymentID uint, items []models.CartItem) (*InvoiceResponse, error) {
	invoiceItems := make([]InvoiceItemRequest, len(items))
	for i, item := range items {
		invoiceItems[i] = InvoiceItemRequest{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	invoiceReq := InvoiceRequest{
		UserID:    userID,
		Total:     total,
		PaymentID: paymentID,
		Items:     invoiceItems,
	}
	jsonData, _ := json.Marshal(invoiceReq)

	req, _ := http.NewRequest("POST", s.config.InvoiceServiceURL+"/invoices", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("invoice service returned status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data InvoiceResponse `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}
