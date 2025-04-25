package controllers

import (
	"fmt" // エラーメッセージを表示するためのパッケージ

	"github.com/stein579/go-weather/weather-app-backend/services" // サービスを読み込むためのパッケージ
	"github.com/stein579/go-weather/weather-app-backend/views"    // ビューを読み込むためのパッケージ
)

type WeatherController struct { // コントローラーを定義する
	service *services.WeatherService // サービスを読み込むためのパッケージ
	view    *views.WeatherView       // ビューを読み込むためのパッケージ
}

func NewWeatherController(apiKey string) *WeatherController { // コントローラーを作成する
	return &WeatherController{ // コントローラーを作成する
		service: services.NewWeatherService(apiKey), // サービスを作成する
		view:    views.NewWeatherView(),             // ビューを作成する
	}
}

func (c *WeatherController) GetCurrentWeather(city string) error { // 現在の天気を取得する
	location, err := c.service.GetLocation(city) // 位置を取得する
	if err != nil {                              // エラーが発生した場合
		return fmt.Errorf("location error: %w", err) // エラーメッセージを表示する
	}

	weather, err := c.service.GetCurrentWeather(location) // 天気を取得する
	if err != nil {                                       // エラーが発生した場合
		return fmt.Errorf("weather error: %w", err) // エラーメッセージを表示する
	}

	output := c.view.FormatCurrentWeather(weather) // 天気を表示する
	fmt.Println(output)                            // 天気を表示する
	return nil                                     // エラーが発生しなかった場合はnilを返す
}

func (c *WeatherController) GetWeatherForecast(city string) error { // 天気予報を取得する
	location, err := c.service.GetLocation(city) // 位置を取得する
	if err != nil {                              // エラーが発生した場合
		return fmt.Errorf("location error: %w", err) // エラーメッセージを表示する
	}

	forecast, err := c.service.GetWeatherForecast(location) // 天気予報を取得する
	if err != nil {                                         // エラーが発生した場合
		return fmt.Errorf("forecast error: %w", err) // エラーメッセージを表示する
	}

	output := c.view.FormatWeatherForecast(forecast) // 天気予報を表示する
	fmt.Println(output)                              // 天気予報を表示する
	return nil                                       // エラーが発生しなかった場合はnilを返す
}
