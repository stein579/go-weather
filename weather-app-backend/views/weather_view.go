package views

import (
	"fmt" // フォーマットを表示するためのパッケージ

	"github.com/stein579/go-weather/weather-app-backend/models" // モデルを読み込むためのパッケージ
)

type WeatherView struct{} // 天気ビューを定義する

func NewWeatherView() *WeatherView { // 天気ビューを作成する
	return &WeatherView{} // 天気ビューを作成する
}

func (v *WeatherView) FormatCurrentWeather(weather *models.Weather) string { // 現在の天気をフォーマットする
	return fmt.Sprintf("天気: %s\n気温: %.2f ℃\n湿度: %.2f %%\n風速: %.2f m/s", // フォーマットを表示する
		weather.Description, // 天気の説明
		weather.Temperature, // 温度
		weather.Humidity,    // 湿度
		weather.WindSpeed,   // 風速
	)
}

func (v *WeatherView) FormatWeatherForecast(forecasts []models.WeatherForecast) string { // 天気予報をフォーマットする
	str := "3時間ごとの天気予報\n"         // フォーマットを表示する
	for _, f := range forecasts { // 天気予報を繰り返す
		str += fmt.Sprintf("%s  %s %.2f ℃\n", // フォーマットを表示する
			f.DateTime,    // 日時
			f.Description, // 天気の説明
			f.Temperature, // 温度
		)
	}
	return str // フォーマットを表示する
}
