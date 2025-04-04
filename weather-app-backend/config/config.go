package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetAPIKey() (string, error) {
	_ = godotenv.Load()
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")

	if apiKey == "" {
		return "", fmt.Errorf("OPEN_WEATHER_API_KEY is not set")
	}

	return apiKey, nil
}
