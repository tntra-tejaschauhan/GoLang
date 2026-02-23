package config

import (
	"os"
)

type Config struct {
	Port                string
	ProductServiceURL   string
	CartServiceURL      string
	InvoiceServiceURL   string
}

func Load() *Config {
	return &Config{
		Port:                getEnv("PORT", "8080"),
		ProductServiceURL:   getEnv("PRODUCT_SERVICE_URL", "http://localhost:8081"),
		CartServiceURL:      getEnv("CART_SERVICE_URL", "http://localhost:8082"),
		InvoiceServiceURL:   getEnv("INVOICE_SERVICE_URL", "http://localhost:8083"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
