package main

import (
	"fmt"
	"os"

	"github.com/stein579/go-weather/weather-app-backend/api"
	"github.com/stein579/go-weather/weather-app-backend/config"
)

func main() {
	apiKey, err := config.GetAPIKey()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	command := os.Args[2]
	city := os.Args[3]

	lat, lon, err := api.GetCoordinates(city, apiKey)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("Coordinates for %s: %f, %f\n", city, lat, lon)

	switch command {
	case "current":
		fmt.Printf("Fetching current weather for %s\n", city)
		// ここに現在の天気を取得する処理を追加
	case "forecast":
		fmt.Printf("Fetching weather forecast for %s\n", city)
		// ここに天気予報を取得する処理を追加
	default:
		fmt.Println("Invalid command. Use 'current' or 'forecast'.")
		os.Exit(1)
	}
}
