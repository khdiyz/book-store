package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	var cfg Config

	cfg.ServerHost = cast.ToString(getValueOrDefault("SERVER_HOST", "localhost"))
	cfg.ServerPort = cast.ToString(getValueOrDefault("SERVER_PORT", "9000"))

	cfg.SQLiteFilePath = cast.ToString(getValueOrDefault("DB_FILE_PATH", "./database/bookstore.db"))

	cfg.SecretKey = cast.ToString(getValueOrDefault("SECRET_KEY", "q6T6LlwdRk"))

	return &cfg
}

type Config struct {
	SQLiteFilePath string

	ServerHost string
	ServerPort string

	SecretKey string
}

func getValueOrDefault(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
