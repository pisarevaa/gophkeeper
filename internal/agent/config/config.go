package config

type Config struct {
	ServerHost string
}

// Инициализация переменных окружения.
func NewConfig() Config {
	config := Config{
		ServerHost: "localhost:8080",
	}
	return config
}
