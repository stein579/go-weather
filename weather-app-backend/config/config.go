package config

import (
	"fmt" // エラーメッセージを表示するためのパッケージ
	"os"  // 環境変数を読み込むためのパッケージ

	"github.com/joho/godotenv" // 環境変数を読み込むためのパッケージ Zenn参考
)

func GetAPIKey() (string, error) {
	_ = godotenv.Load()                         // .envファイルから環境変数を読み込む
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY") // 環境変数からAPIキーを取得する

	if apiKey == "" {
		return "", fmt.Errorf("OPEN_WEATHER_API_KEY is not set") // APIキーが設定されていない(空文字)場合はエラーを返す
	}

	return apiKey, nil // APIキーを返す
}
