package config

import "os"

type Config struct {
	BotToken    string
	Port        string
	DatabaseURL string
}

func LoadConfig() *Config {
	return &Config{
		BotToken:    getEnvOrDefault("BOT_TOKEN", ""),
		Port:        getEnvOrDefault("PORT", "8080"),
		DatabaseURL: getEnvOrDefault("DATABASE_URL", ""),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
