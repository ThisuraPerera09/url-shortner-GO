package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port         string
	BaseURL      string
	DatabasePath string
	ShortCodeLen int
	UseInMemory  bool
}

func Load() *Config {
	baseURL := getEnv("BASE_URL", "http://localhost:8080")
	// Remove trailing slash to avoid double slashes in URLs
	baseURL = strings.TrimSuffix(baseURL, "/")

	return &Config{
		Port:         getEnv("PORT", "8080"),
		BaseURL:      baseURL,
		DatabasePath: getEnv("DATABASE_PATH", "./urlshortener.db"),
		ShortCodeLen: getEnvAsInt("SHORT_CODE_LEN", 6),
		UseInMemory:  getEnvAsBool("USE_IN_MEMORY", true), // Changed default to true
	}
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultVal
}
