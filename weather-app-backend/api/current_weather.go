package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CurrentWeatherResponse struct {
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

func GetCurrentWeather(lat float64, lon float64, apiKey string) (string, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lang=ja&units=metric&lat=%f&lon=%f&appid=%s", lat, lon, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api error: status code %d", resp.StatusCode)
	}

	var result CurrentWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	str := fmt.Sprintf("天気: %s\n気温: %.2f ℃\n湿度: %.2f %%\n風速: %.2f m/s",
		result.Weather[0].Description, result.Main.Temp, result.Main.Humidity, result.Wind.Speed)
	return str, nil
}
