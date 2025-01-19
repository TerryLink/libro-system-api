package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName                string
	Port                   string
	JWTExpirationInSeconds int64
	JWTSecret              string
	Database               struct {
		User     string
		Password string
		Name     string
		Host     string
		Port     string
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	cfg := &Config{
		AppName:                getEnv("APP_NAME", "local"),
		Port:                   getEnv("PORT", "8080"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "default-secret"),
		Database: struct {
			User     string
			Password string
			Name     string
			Host     string
			Port     string
		}{
			User:     getEnv("BLUEPRINT_DB_USERNAME", "melkey"),
			Password: getEnv("BLUEPRINT_DB_PASSWORD", "password1234"),
			Name:     getEnv("BLUEPRINT_DB_DATABASE", "blueprint"),
			Host:     getEnv("BLUEPRINT_DB_PORT", "3306"),
			Port:     getEnv("BLUEPRINT_DB_HOST", "localhost"),
		},
	}
	return cfg, nil
}
