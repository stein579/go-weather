package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GeocodeResponse struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func GetCoordinates(city string, apiKey string) (float64, float64, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("api error: status code %d", resp.StatusCode)
	}

	var result []GeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, err
	}

	if len(result) == 0 {
		return 0, 0, fmt.Errorf("no results")
	}

	return result[0].Lat, result[0].Lon, nil
}
