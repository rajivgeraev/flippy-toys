package config

import "os"

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
	return &Config{
		BotToken:               getEnvOrDefault("BOT_TOKEN", ""),
		Port:                   getEnvOrDefault("PORT", "8080"),
		DatabaseURL:            getEnvOrDefault("DATABASE_URL", ""),
		CloudinaryName:         getEnvOrDefault("CLOUDINARY_CLOUD_NAME", ""),
		CloudinaryApiKey:       getEnvOrDefault("CLOUDINARY_API_KEY", ""),
		CloudinaryApiSecret:    getEnvOrDefault("CLOUDINARY_API_SECRET", ""),
		CloudinaryUploadPreset: getEnvOrDefault("CLOUDINARY_UPLOAD_PRESET", ""),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
