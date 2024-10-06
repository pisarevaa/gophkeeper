package utils

import (
	"encoding/json"
	"os"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

const UserFilename = "gophkeeper_auth.json"

// Сохранение данных пользователя на диск.
func SaveUserDataToDisk(tokenResponse model.TokenResponse) error {
	file, err := os.OpenFile(UserFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(&tokenResponse)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

// Получение данных пользователя с диска.
func LoadUserDataFromDosk() (model.TokenResponse, error) {
	file, err := os.OpenFile(UserFilename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return model.TokenResponse{}, err
	}
	var tokenResponse model.TokenResponse
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tokenResponse)
	if err != nil {
		return model.TokenResponse{}, err
	}
	err = file.Close()
	if err != nil {
		return model.TokenResponse{}, err
	}
	return tokenResponse, nil
}
