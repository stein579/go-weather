package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/stein579/go-weather/weather-app-backend/models"
)

type WeatherService struct {
	apiKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{apiKey: apiKey}
}

func (s *WeatherService) GetLocation(city string) (*models.Location, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", city, s.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("city not found")
	}

	return &models.Location{
		City:      city,
		Latitude:  result[0].Lat,
		Longitude: result[0].Lon,
	}, nil
}

func (s *WeatherService) GetCurrentWeather(location *models.Location) (*models.Weather, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lang=ja&units=metric&lat=%f&lon=%f&appid=%s",
		location.Latitude, location.Longitude, s.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity float64 `json:"humidity"`
		} `json:"main"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &models.Weather{
		Description: result.Weather[0].Description,
		Temperature: result.Main.Temp,
		Humidity:    result.Main.Humidity,
		WindSpeed:   result.Wind.Speed,
	}, nil
}

func (s *WeatherService) GetWeatherForecast(location *models.Location) ([]models.WeatherForecast, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lang=ja&units=metric&lat=%f&lon=%f&appid=%s",
		location.Latitude, location.Longitude, s.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		List []struct {
			DtTxt   string `json:"dt_txt"`
			Weather []struct {
				Description string `json:"description"`
			} `json:"weather"`
			Main struct {
				Temp float64 `json:"temp"`
			} `json:"main"`
		} `json:"list"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	forecasts := make([]models.WeatherForecast, 0, len(result.List))
	jst, _ := time.LoadLocation("Asia/Tokyo")

	for _, item := range result.List {
		t, _ := time.Parse("2006-01-02 15:04:05", item.DtTxt)
		t = t.In(jst)
		formattedTime := fmt.Sprintf("%d月%d日 %02d:%02d", t.Month(), t.Day(), t.Hour(), t.Minute())

		forecasts = append(forecasts, models.WeatherForecast{
			DateTime:    formattedTime,
			Description: item.Weather[0].Description,
			Temperature: item.Main.Temp,
		})
	}

	return forecasts, nil
}
