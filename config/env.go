package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv は環境変数を読み込む
// プロファイルに応じた.envファイルを読み込む
func LoadEnv(profile string) error {

	var envFile string = profile + ".env"

	// 環境変数をロードする
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading .env file for profile %s", profile)
		return err
	}

	return nil
}

// getEnv は環境変数を取得するユーティリティ関数（デフォルト値も設定可能）
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
