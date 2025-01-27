package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken               string
	Port                   string
	DatabaseURL            string
	CloudinaryName         string
	CloudinaryApiKey       string
	CloudinaryApiSecret    string
	CloudinaryUploadPreset string
}

func LoadConfig() *Config {
	// Try to load .env file first
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found or cannot be loaded: %v", err)
	}

	// Load configuration from environment variables (either from .env or system env)
	config := &Config{
		BotToken:               getEnvOrDefault("BOT_TOKEN", ""),
		Port:                   getEnvOrDefault("PORT", "8080"),
		DatabaseURL:            getEnvOrDefault("DATABASE_URL", ""),
		CloudinaryName:         getEnvOrDefault("CLOUDINARY_CLOUD_NAME", ""),
		CloudinaryApiKey:       getEnvOrDefault("CLOUDINARY_API_KEY", ""),
		CloudinaryApiSecret:    getEnvOrDefault("CLOUDINARY_API_SECRET", ""),
		CloudinaryUploadPreset: getEnvOrDefault("CLOUDINARY_UPLOAD_PRESET", ""),
	}

	return config
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
