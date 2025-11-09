package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	DatabaseURL string
	JWTSecret   string
	GinMode     string
	Port        string
}

// Load returns application configuration from environment variables
func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://localhost/asstbackend?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
		GinMode:     getEnv("GIN_MODE", "debug"),
		Port:        getEnv("PORT", "8080"),
	}
}

// getEnv returns environment variable value or default if not set
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
