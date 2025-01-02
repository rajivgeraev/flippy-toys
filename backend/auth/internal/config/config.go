package config

import (
    "fmt"
    "os"
)

type Config struct {
    PostgresConfig PostgresConfig
    ServerConfig  ServerConfig
}

type PostgresConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

type ServerConfig struct {
    Port string
}

func New() (*Config, error) {
    postgresConfig := PostgresConfig{
        Host:     getEnv("POSTGRES_HOST", "localhost"),
        Port:     getEnv("POSTGRES_PORT", "5432"),
        User:     getEnv("POSTGRES_USER", "postgres"),
        Password: getEnv("POSTGRES_PASSWORD", ""),
        DBName:   getEnv("POSTGRES_DB", "flippy"),
        SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
    }

    serverConfig := ServerConfig{
        Port: getEnv("PORT", "8080"),
    }

    return &Config{
        PostgresConfig: postgresConfig,
        ServerConfig:  serverConfig,
    }, nil
}

func (c *PostgresConfig) GetDSN() string {
    return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
    )
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}