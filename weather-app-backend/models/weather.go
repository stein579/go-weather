package models

type Weather struct {
	Description string
	Temperature float64
	Humidity    float64
	WindSpeed   float64
}

type WeatherForecast struct {
	DateTime    string
	Description string
	Temperature float64
}

type WeatherData struct {
	Current  *Weather
	Forecast []WeatherForecast
}
