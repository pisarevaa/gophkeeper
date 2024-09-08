package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Database struct {
	Dsn    string
	Schema string
}

type Kafka struct {
	Server    string
	TopicGeo  string
	TopicStat string
}

type Config struct {
	Host     string
	Database Database
	Kafka    Kafka
}

func (s *Config) SetHost() {
	s.Host = viper.GetString("GO_HOST")
	if s.Host == "" {
		s.Host = "localhost"
	}
}

func (s *Config) SetDatabase() {
	s.Database.Dsn = viper.GetString("DB_DSN")
	if s.Database.Dsn == "" {
		s.Database.Dsn = "postgresql://admin:CC7B02B06C4C1CF81FAE7D8C46C429EC@localhost:5432/sdk-backend"
	}
	s.Database.Schema = viper.GetString("DB_SCHEMA")
	if s.Database.Schema == "" {
		s.Database.Schema = "backend"
	}
}

func (s *Config) SetKafka() {
	s.Kafka.Server = viper.GetString("KAFKA_BOOTSTRAP_SERVER")
	if s.Kafka.Server == "" {
		s.Kafka.Server = "localhost:9092"
	}
	s.Kafka.TopicStat = viper.GetString("KAFKA_TOPIC")
	if s.Kafka.Server == "" {
		s.Kafka.Server = "app_logs"
	}
	s.Kafka.TopicGeo = viper.GetString("KAFKA_TOPIC_GEO")
	if s.Kafka.TopicGeo == "" {
		s.Kafka.TopicGeo = "geo_logs"
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
	config.SetKafka()
	return config
}
