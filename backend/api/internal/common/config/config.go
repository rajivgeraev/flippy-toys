package config

import "os"

type Config struct {
	BotToken    string
	Port        string
	DatabaseURL string

	CloudinaryName   string
	CloudinaryKey    string
	CloudinarySecret string
}

func LoadConfig() *Config {
	return &Config{
		BotToken:    getEnvOrDefault("BOT_TOKEN", ""),
		Port:        getEnvOrDefault("PORT", "8080"),
		DatabaseURL: getEnvOrDefault("DATABASE_URL", ""),

		CloudinaryName:   getEnvOrDefault("CLOUDINARY_CLOUD_NAME", ""),
		CloudinaryKey:    getEnvOrDefault("CLOUDINARY_API_KEY", ""),
		CloudinarySecret: getEnvOrDefault("CLOUDINARY_API_SECRET", ""),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
