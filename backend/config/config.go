package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	JWTSecret   string
	Environment string
	DATABASE_URL string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	config := &Config{
		JWTSecret:   os.Getenv("JWT_SECRET"),
		Environment: os.Getenv("ENV"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}
	return config
}
