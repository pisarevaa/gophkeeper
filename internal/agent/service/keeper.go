package service

import (
	"encoding/json"
	"log/slog"

	"github.com/pisarevaa/gophkeeper/internal/agent/utils"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Получение всех данных пользователя.
func (s *Service) GetData() error {
	if err := s.Client.SetToken(); err != nil {
		return err
	}
	dataResponse, err := s.Client.GetData()
	if err != nil {
		return err
	}
	slog.Info("List of data:")
	for _, row := range dataResponse {
		rowString, errJSON := json.Marshal(row)
		if errJSON != nil {
			return errJSON
		}
		slog.Info(string(rowString))
	}
	return nil
}

// Получение данных пользователя по ID.
func (s *Service) GetDataByID(dataID int64) error {
	if err := s.Client.SetToken(); err != nil {
		return err
	}
	dataResponse, err := s.Client.GetDataByID(dataID)
	if err != nil {
		return err
	}
	if dataResponse.Type == model.TextType {
		decryptedData, errDecrypt := utils.DecryptString(s.Config.PrivateKey, dataResponse.Data)
		if errDecrypt != nil {
			return errDecrypt
		}
		dataResponse.Data = decryptedData
	}
	rowString, err := json.Marshal(dataResponse)
	if err != nil {
		return err
	}
	slog.Info("Get data:")
	slog.Info(string(rowString))
	return nil
}

// Добавление текстовых данных пользователем.
func (s *Service) AddTextData(textData model.AddTextData) error {
	if err := s.Client.SetToken(); err != nil {
		return err
	}
	encryptedData, err := utils.EncryptString(s.Config.PublicKey, []byte(textData.Data))
	if err != nil {
		return err
	}
	textData.Data = encryptedData
	dataResponse, err := s.Client.AddTextData(textData)
	if err != nil {
		return err
	}
	rowString, err := json.Marshal(dataResponse)
	if err != nil {
		return err
	}
	slog.Info("Add new text data:")
	slog.Info(string(rowString))
	return nil
}

// Обновление текстовых данных пользователем.
func (s *Service) UpdateTextData(textData model.AddTextData, dataID int64) error {
	if err := s.Client.SetToken(); err != nil {
		return err
	}
	encryptedData, err := utils.EncryptString(s.Config.PublicKey, []byte(textData.Data))
	if err != nil {
		return err
	}
	textData.Data = encryptedData
	dataResponse, err := s.Client.UpdateTextData(textData, dataID)
	if err != nil {
		return err
	}
	rowString, err := json.Marshal(dataResponse)
	if err != nil {
		return err
	}
	slog.Info("Updated text data:")
	slog.Info(string(rowString))
	return nil
}

// Добавление бинарных данных пользователем.
func (s *Service) AddBinaryData(filepath string, name string) error {
	if err := s.Client.SetToken(); err != nil {
		return err
	}
	dataResponse, err := s.Client.AddBinaryData(filepath, name)
	if err != nil {
		return err
	}
	rowString, err := json.Marshal(dataResponse)
	if err != nil {
		return err
	}
	slog.Info("Add new binary data:")
	slog.Info(string(rowString))
	return nil
}

// Обновление бинарных данных пользователем.
func (s *Service) UpdateBinaryData(filepath string, name string, dataID int64) error {
	if err := s.Client.SetToken(); err != nil {
		return err
	}
	dataResponse, err := s.Client.UpdateBinaryData(filepath, name, dataID)
	if err != nil {
		return err
	}
	rowString, err := json.Marshal(dataResponse)
	if err != nil {
		return err
	}
	slog.Info("Updated binary data:")
	slog.Info(string(rowString))
	return nil
}

// Удаление данных пользователя по ID.
func (s *Service) DeleteData(dataID int64) error {
	if err := s.Client.SetToken(); err != nil {
		return err
	}
	dataResponse, err := s.Client.DeleteData(dataID)
	if err != nil {
		return err
	}
	rowString, err := json.Marshal(dataResponse)
	if err != nil {
		return err
	}
	slog.Info("Deleted data:")
	slog.Info(string(rowString))
	return nil
}
