package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Database struct {
	Dsn    string
	Schema string
}

type Config struct {
	Host     string
	Database Database
}

func (s *Config) SetHost() {
	s.Host = viper.GetString("HOST")
	if s.Host == "" {
		s.Host = "localhost:5000"
	}
}

func (s *Config) SetDatabase() {
	s.Database.Dsn = viper.GetString("DB_DSN")
	if s.Database.Dsn == "" {
		s.Database.Dsn = "postgresql://gophkeeper:CC7B02B06C4C1CF81FAE7D8C46C429EC@localhost:5432/gophkeeper"
	}
	s.Database.Schema = viper.GetString("DB_SCHEMA")
	if s.Database.Schema == "" {
		s.Database.Schema = "gophkeeper"
	}
}

// Инициализация переменных окружения.
func NewConfig(logger *zap.SugaredLogger) Config {
	var config Config
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err)
		viper.AutomaticEnv()
	}
	config.SetHost()
	config.SetDatabase()
	return config
}
