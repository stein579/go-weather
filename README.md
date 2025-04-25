## 概要
Go言語で作成された天気予報アプリケーションです。OpenWeather APIを使用して、指定した都市の現在の天気と3時間毎の天気予報を取得できます。MVCアーキテクチャを採用し、コードの保守性と拡張性を高めています。

## 使用技術
- Go 1.24.0
- OpenWeather API
- godotenv（環境変数管理）

## セットアップ手順
1. リポジトリをクローン
```bash
git clone https://github.com/stein579/go-weather.git
cd go-weather
```

2. 依存関係のインストール
```bash
go mod tidy
```

3. 環境変数の設定
`.env`ファイルを作成し、OpenWeather APIのキーを設定：
```
OPEN_WEATHER_API_KEY=your_api_key_here
```

4. アプリケーションの実行
現在の天気を取得：
```bash
go run weather-app-backend/main.go api current "Tokyo"
```

天気予報を取得：
```bash
go run weather-app-backend/main.go api forecast "Tokyo"
```

## ディレクトリ構成
```
weather-app-backend/
├── models/          # データモデルの定義
│   ├── weather.go   # 天気情報のモデル
│   └── location.go  # 位置情報のモデル
├── views/           # 表示関連の処理
│   └── weather_view.go
├── controllers/     # ビジネスロジックの制御
│   └── weather_controller.go
├── services/        # APIとの通信処理
│   └── weather_service.go
├── config/          # 設定関連
│   └── config.go
└── main.go         # エントリーポイント
```

## 機能一覧
- 現在の天気情報の取得
  - 天気の状態
  - 気温
  - 湿度
  - 風速
- 3時間ごとの天気予報の取得
  - 日時
  - 天気の状態
  - 気温

## 今後の改善点
- フロントで都市を選択、検索して天気を表示する
- エラーハンドリングの強化
- ユニットテストの追加
- キャッシュ機能の実装
- より詳細な天気情報の表示
- Webインターフェースの追加
- 複数都市の同時検索機能


