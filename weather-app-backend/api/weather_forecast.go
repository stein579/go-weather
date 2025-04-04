package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherForecastResponse struct {
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

func GetWeatherForecast(lat float64, lon float64, apiKey string) (string, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lang=ja&units=metric&lat=%f&lon=%f&appid=%s", lat, lon, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api error: status code %d", resp.StatusCode)
	}

	var result WeatherForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	str := "3時間ごとの天気予報\n"

	jst, _ := time.LoadLocation("Asia/Tokyo")

	for _, forecast := range result.List {
		// 日付フォーマットを整形
		t, _ := time.Parse("2006-01-02 15:04:05", forecast.DtTxt) // UTC でパース
		t = t.In(jst)                                             // JST に変換
		formattedTime := fmt.Sprintf("%d月%d日 %02d:%02d", t.Month(), t.Day(), t.Hour(), t.Minute())

		str += fmt.Sprintf("%s  %s %.2f ℃\n", formattedTime, forecast.Weather[0].Description, forecast.Main.Temp)
	}

	return str, nil
}
