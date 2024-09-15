package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	Dsn string
}

type Security struct {
	SecretKey   string
	TokenExpSec int64
}

type Minio struct {
	User     string
	Password string
	UseSSL   bool
	Bucket   string
}

type Config struct {
	Host     string
	Security Security
	Database Database
	Minio    Minio
}

// Инициализация переменных окружения.
func NewConfig() Config {
	var config Config
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
	}
	config = Config{
		Host: getEnvString("HOST", "localhost:8080"),
		Security: Security{
			SecretKey:   getEnvString("SECURITY_SECRET_KEY", "7fd315fd5f381bb9035d003dbd904102"),
			TokenExpSec: getEnvInt64("SECURITY_TOKEN_EXP_SEC", 7200),
		},
		Database: Database{
			Dsn: getEnvString(
				"DB_DSN",
				"postgresql://gophkeeper:CC7B02B06C4C1CF81FAE7D8C46C429EC@localhost:5432/gophkeeper?sslmode=disable",
			),
		},
		Minio: Minio{
			User:     getEnvString("MINIO_ROOT_USER", "gophkeeper"),
			Password: getEnvString("MINIO_ROOT_PASSWORD", "gophkeeper"),
			UseSSL:   getEnvBool("MINIO_USE_SSL", false),
			Bucket:   getEnvString("MINIO_DEFAULT_BUCKETS", "gophkeeper"),
		},
	}
	return config
}

func getEnvString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt64(key string, defaultValue int64) int64 {
	if valueStr := getEnvString(key, ""); valueStr != "" {
		if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
			return value
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if valueStr := getEnvString(key, ""); valueStr != "" {
		if value, err := strconv.ParseBool(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}
