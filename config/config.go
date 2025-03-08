package config

// Config はアプリケーションの設定を保持する構造体
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// LoadConfig は環境変数を使って設定をロードし、Config構造体に格納する
func LoadConfig(profile string) (*Config, error) {
	// 環境変数をロード
	err := LoadEnv(profile)
	if err != nil {
		return nil, err
	}

	// 設定を Config 構造体に格納
	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "waritomodb"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return config, nil
}
