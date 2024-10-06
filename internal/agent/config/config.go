package config

import (
	"crypto/rsa"

	"github.com/pisarevaa/gophkeeper/internal/agent/utils"
)

const (
	publicKeyFilepath  = "gophkeeper_public.key"
	privateKeyFilepath = "gophkeeper_private.key"
)

type Config struct {
	ServerHost string
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// Инициализация переменных окружения.
func NewConfig() (*Config, error) {
	publicKey, err := utils.InitPublicKey(publicKeyFilepath)
	if err != nil {
		return nil, err
	}
	privateKey, err := utils.InitPrivateKey(privateKeyFilepath)
	if err != nil {
		return nil, err
	}
	config := Config{
		ServerHost: "http://127.0.0.1:8080",
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
	return &config, nil
}
