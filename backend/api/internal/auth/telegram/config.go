package telegram

import (
	"os"
)

// DevModeConfig holds development mode configuration
type DevModeConfig struct {
	Enabled   bool
	DevUserID int64
}

// GetDevModeConfig returns the development mode configuration
func GetDevModeConfig() *DevModeConfig {
	devMode := os.Getenv("DEV_MODE") == "true"

	// Only enable dev mode if explicitly set and not in production
	if os.Getenv("ENV") == "production" {
		devMode = false
	}

	// Parse dev user ID if provided
	var devUserID int64 = 136833584 // Default to the provided test user ID

	return &DevModeConfig{
		Enabled:   devMode,
		DevUserID: devUserID,
	}
}
