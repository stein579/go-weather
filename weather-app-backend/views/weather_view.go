package views

import (
	"fmt"

	"github.com/stein579/go-weather/weather-app-backend/models"
)

type WeatherView struct{}

func NewWeatherView() *WeatherView {
	return &WeatherView{}
}

func (v *WeatherView) FormatCurrentWeather(weather *models.Weather) string {
	return fmt.Sprintf("天気: %s\n気温: %.2f ℃\n湿度: %.2f %%\n風速: %.2f m/s",
		weather.Description,
		weather.Temperature,
		weather.Humidity,
		weather.WindSpeed,
	)
}

func (v *WeatherView) FormatWeatherForecast(forecasts []models.WeatherForecast) string {
	str := "3時間ごとの天気予報\n"
	for _, f := range forecasts {
		str += fmt.Sprintf("%s  %s %.2f ℃\n",
			f.DateTime,
			f.Description,
			f.Temperature,
		)
	}
	return str
}
