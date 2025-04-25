package services

import (
	"encoding/json" // JSONをデコードするためのパッケージ
	"fmt"           // エラーメッセージを表示するためのパッケージ
	"net/http"      // HTTPリクエストを送信するためのパッケージ
	"time"          // 時間を扱うためのパッケージ

	"github.com/stein579/go-weather/weather-app-backend/models" // モデルを読み込むためのパッケージ
)

type WeatherService struct { // 天気サービスを定義する
	apiKey string // APIキー
}

func NewWeatherService(apiKey string) *WeatherService { // 天気サービスを作成する
	return &WeatherService{apiKey: apiKey} // 天気サービスを作成する
}

func (s *WeatherService) GetLocation(city string) (*models.Location, error) { // 位置を取得する
	url := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", city, s.apiKey) // URLを作成する
	resp, err := http.Get(url)                                                                                // HTTPリクエストを送信する
	if err != nil {                                                                                           // エラーが発生した場合
		return nil, err // エラーを返す
	}
	defer resp.Body.Close() // レスポンスを閉じる

	var result []struct { // 位置を格納する
		Lat float64 `json:"lat"` // 緯度
		Lon float64 `json:"lon"` // 経度
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // JSONをデコードする
		return nil, err // エラーを返す
	}

	if len(result) == 0 { // 位置が見つからない場合
		return nil, fmt.Errorf("city not found") // エラーを返す
	}

	return &models.Location{ // 位置を返す
		City:      city,          // 都市名
		Latitude:  result[0].Lat, // 緯度
		Longitude: result[0].Lon, // 経度
	}, nil // エラーが発生しなかった場合はnilを返す
}

func (s *WeatherService) GetCurrentWeather(location *models.Location) (*models.Weather, error) { // 現在の天気を取得する
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lang=ja&units=metric&lat=%f&lon=%f&appid=%s", // URLを作成する
		location.Latitude, location.Longitude, s.apiKey) // 緯度と経度をURLに追加する

	resp, err := http.Get(url) // HTTPリクエストを送信する
	if err != nil {            // エラーが発生した場合
		return nil, err // エラーを返す
	}
	defer resp.Body.Close() // レスポンスを閉じる

	var result struct { // 天気を格納する
		Weather []struct { // 天気を格納する
			Description string `json:"description"` // 天気の説明
		} `json:"weather"` // 天気を格納する
		Main struct { // 天気を格納する
			Temp     float64 `json:"temp"`     // 温度
			Humidity float64 `json:"humidity"` // 湿度
		} `json:"main"` // 天気を格納する
		Wind struct { // 天気を格納する
			Speed float64 `json:"speed"` // 風速
		} `json:"wind"` // 天気を格納する
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // JSONをデコードする
		return nil, err // エラーを返す
	}

	return &models.Weather{ // 天気を返す
		Description: result.Weather[0].Description, // 天気の説明
		Temperature: result.Main.Temp,              // 温度
		Humidity:    result.Main.Humidity,          // 湿度
		WindSpeed:   result.Wind.Speed,             // 風速
	}, nil // エラーが発生しなかった場合はnilを返す
}

func (s *WeatherService) GetWeatherForecast(location *models.Location) ([]models.WeatherForecast, error) { // 天気予報を取得する
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lang=ja&units=metric&lat=%f&lon=%f&appid=%s", // URLを作成する
		location.Latitude, location.Longitude, s.apiKey) // 緯度と経度をURLに追加する

	resp, err := http.Get(url) // HTTPリクエストを送信する
	if err != nil {            // エラーが発生した場合
		return nil, err // エラーを返す
	}
	defer resp.Body.Close() // レスポンスを閉じる

	var result struct { // 天気予報を格納する
		List []struct { // 天気予報を格納する
			DtTxt   string     `json:"dt_txt"` // 日時
			Weather []struct { // 天気予報を格納する
				Description string `json:"description"` // 天気の説明
			} `json:"weather"` // 天気予報を格納する
			Main struct { // 天気予報を格納する
				Temp float64 `json:"temp"` // 温度
			} `json:"main"` // 天気予報を格納する
		} `json:"list"` // 天気予報を格納する
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // JSONをデコードする
		return nil, err // エラーを返す
	}

	forecasts := make([]models.WeatherForecast, 0, len(result.List)) // 天気予報を格納する
	jst, _ := time.LoadLocation("Asia/Tokyo")                        // 東京のタイムゾーンを読み込む

	for _, item := range result.List { // 天気予報を繰り返す
		t, _ := time.Parse("2006-01-02 15:04:05", item.DtTxt)                                      // 日時を解析する
		t = t.In(jst)                                                                              // 東京のタイムゾーンに変換する
		formattedTime := fmt.Sprintf("%d月%d日 %02d:%02d", t.Month(), t.Day(), t.Hour(), t.Minute()) // 日時をフォーマットする

		forecasts = append(forecasts, models.WeatherForecast{ // 天気予報を追加する
			DateTime:    formattedTime,               // 日時
			Description: item.Weather[0].Description, // 天気の説明
			Temperature: item.Main.Temp,              // 温度
		})
	}

	return forecasts, nil // エラーが発生しなかった場合はnilを返す
}
