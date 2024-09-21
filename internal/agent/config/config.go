package config

type Config struct {
	ServerHost string
}

// Инициализация переменных окружения.
func NewConfig() Config {
	config := Config{
		ServerHost: "http://127.0.0.1:8080",
	}
	return config
}
