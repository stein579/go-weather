package main

import (
	"fmt" // エラーメッセージを表示するためのパッケージ
	"os"  // 環境変数を読み込むためのパッケージ

	"github.com/stein579/go-weather/weather-app-backend/config"      // 設定ファイルを読み込むためのパッケージ
	"github.com/stein579/go-weather/weather-app-backend/controllers" // コントローラーを読み込むためのパッケージ
)

func main() {
	apiKey, err := config.GetAPIKey() // APIキーを取得する
	if err != nil {
		fmt.Println("Error: ", err) // エラーメッセージを表示する
		os.Exit(1)                  // プログラムを終了する、Qiita参考
	}

	controller := controllers.NewWeatherController(apiKey) // コントローラーを作成する

	if len(os.Args) != 4 { // コマンドライン引数の数が4つでない場合
		fmt.Println("Usage: go run main.go api [current|forecast] CITY") // コマンドライン引数の数が4つでない場合はエラーメッセージを表示する
		os.Exit(1)                                                       // プログラムを終了する、Qiita参考
	}

	command := os.Args[2] // コマンドライン引数の2番目の要素を取得する
	city := os.Args[3]    // コマンドライン引数の3番目の要素を取得する

	var cmdErr error // コマンドライン引数の数が4つでない場合はエラーメッセージを表示する
	switch command { // コマンドライン引数の2番目の要素を取得する
	case "current": // コマンドライン引数の2番目の要素が"current"の場合
		cmdErr = controller.GetCurrentWeather(city) // コントローラーのGetCurrentWeatherメソッドを呼び出す
	case "forecast": // コマンドライン引数の2番目の要素が"forecast"の場合
		cmdErr = controller.GetWeatherForecast(city) // コントローラーのGetWeatherForecastメソッドを呼び出す
	default: // コマンドライン引数の2番目の要素が"current"も"forecast"もない場合
		fmt.Println("Invalid command. Use 'current' or 'forecast'.") // エラーメッセージを表示する
		os.Exit(1)                                                   // プログラムを終了する、Qiita参考
	}

	if cmdErr != nil { // コマンドライン引数の数が4つでない場合はエラーメッセージを表示する
		fmt.Println("Error:", cmdErr) // エラーメッセージを表示する
		os.Exit(1)                    // プログラムを終了する、Qiita参考
	}
}
