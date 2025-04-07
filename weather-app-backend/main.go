package main

import (
	"fmt"
	"os"

	"github.com/stein579/go-weather/weather-app-backend/config"
	"github.com/stein579/go-weather/weather-app-backend/controllers"
)

func main() {
	apiKey, err := config.GetAPIKey()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	controller := controllers.NewWeatherController(apiKey)

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go api [current|forecast] CITY")
		os.Exit(1)
	}

	command := os.Args[2]
	city := os.Args[3]

	var cmdErr error
	switch command {
	case "current":
		cmdErr = controller.GetCurrentWeather(city)
	case "forecast":
		cmdErr = controller.GetWeatherForecast(city)
	default:
		fmt.Println("Invalid command. Use 'current' or 'forecast'.")
		os.Exit(1)
	}

	if cmdErr != nil {
		fmt.Println("Error:", cmdErr)
		os.Exit(1)
	}
}
