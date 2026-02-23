package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port         string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPass       string
	DBName       string
	PaymentLimit float64  // Amount threshold for payment approval
}

func Load() *Config {
	limitStr := getEnv("PAYMENT_LIMIT", "1000.0")
	limit, _ := strconv.ParseFloat(limitStr, 64)

	return &Config{
		Port:         getEnv("PORT", "8084"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5435"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPass:       getEnv("DB_PASS", "postgres"),
		DBName:       getEnv("DB_NAME", "payment_db"),
		PaymentLimit: limit,
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
