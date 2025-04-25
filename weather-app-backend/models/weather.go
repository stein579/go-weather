package models

type Weather struct { // 天気を定義する
	Description string  // 天気の説明
	Temperature float64 // 温度
	Humidity    float64 // 湿度
	WindSpeed   float64 // 風速
}

type WeatherForecast struct { // 天気予報を定義する
	DateTime    string  // 日時
	Description string  // 天気の説明
	Temperature float64 // 温度
}

type WeatherData struct { // 天気データを定義する
	Current  *Weather          // 現在の天気
	Forecast []WeatherForecast // 天気予報
}
