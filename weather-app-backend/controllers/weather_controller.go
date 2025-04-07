package controllers

import (
	"fmt"

	"github.com/stein579/go-weather/weather-app-backend/services"
	"github.com/stein579/go-weather/weather-app-backend/views"
)

type WeatherController struct {
	service *services.WeatherService
	view    *views.WeatherView
}

func NewWeatherController(apiKey string) *WeatherController {
	return &WeatherController{
		service: services.NewWeatherService(apiKey),
		view:    views.NewWeatherView(),
	}
}

func (c *WeatherController) GetCurrentWeather(city string) error {
	location, err := c.service.GetLocation(city)
	if err != nil {
		return fmt.Errorf("location error: %w", err)
	}

	weather, err := c.service.GetCurrentWeather(location)
	if err != nil {
		return fmt.Errorf("weather error: %w", err)
	}

	output := c.view.FormatCurrentWeather(weather)
	fmt.Println(output)
	return nil
}

func (c *WeatherController) GetWeatherForecast(city string) error {
	location, err := c.service.GetLocation(city)
	if err != nil {
		return fmt.Errorf("location error: %w", err)
	}

	forecast, err := c.service.GetWeatherForecast(location)
	if err != nil {
		return fmt.Errorf("forecast error: %w", err)
	}

	output := c.view.FormatWeatherForecast(forecast)
	fmt.Println(output)
	return nil
}
