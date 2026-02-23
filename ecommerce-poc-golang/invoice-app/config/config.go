package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port                  string
	DBHost                string
	DBPort                string
	DBUser                string
	DBPass                string
	DBName                string
	ManagementServiceURL  string
}

func Load() *Config {
	return &Config{
		Port:                 getEnv("PORT", "8083"),
		DBHost:               getEnv("DB_HOST", "localhost"),
		DBPort:               getEnv("DB_PORT", "5434"),
		DBUser:               getEnv("DB_USER", "postgres"),
		DBPass:               getEnv("DB_PASS", "postgres"),
		DBName:               getEnv("DB_NAME", "invoice_db"),
		ManagementServiceURL: getEnv("MANAGEMENT_SERVICE_URL", "http://localhost:8085"),
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
