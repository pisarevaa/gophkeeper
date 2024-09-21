package service

import (
	"encoding/json"
	"log/slog"

	"github.com/pisarevaa/gophkeeper/internal/agent/utils"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
	sharedUtils "github.com/pisarevaa/gophkeeper/internal/shared/utils"
)

// Получение всех данных пользователя.
func (s *Service) GetData() error {
	tokenResponse, err := utils.LoadUserDataFromDosk()
	if err != nil {
		return err
	}
	s.Client.SetToken(tokenResponse.Token)
	dataResponse, err := s.Client.GetData()
	if err != nil {
		return err
	}
	slog.Info("List of data:")
	for _, row := range dataResponse {
		rowString, err := json.Marshal(row)
		if err != nil {
			return err
		}
		slog.Info(string(rowString))
	}
	return nil
}

// Получение данных пользователя по ID.
func (s *Service) GetDataByID(dataID int64) error {
	tokenResponse, err := utils.LoadUserDataFromDosk()
	if err != nil {
		return err
	}
	s.Client.SetToken(tokenResponse.Token)
	dataResponse, err := s.Client.GetDataByID(dataID)
	if err != nil {
		return err
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
	tokenResponse, err := utils.LoadUserDataFromDosk()
	if err != nil {
		return err
	}
	s.Client.SetToken(tokenResponse.Token)
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
	tokenResponse, err := utils.LoadUserDataFromDosk()
	if err != nil {
		return err
	}
	s.Client.SetToken(tokenResponse.Token)
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
	tokenResponse, err := utils.LoadUserDataFromDosk()
	if err != nil {
		return err
	}
	s.Client.SetToken(tokenResponse.Token)
	formData, err := sharedUtils.CreateFormData(filepath, name)
	if err != nil {
		return err
	}
	dataResponse, err := s.Client.AddBinaryData(formData)
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
	tokenResponse, err := utils.LoadUserDataFromDosk()
	if err != nil {
		return err
	}
	s.Client.SetToken(tokenResponse.Token)
	formData, err := sharedUtils.CreateFormData(filepath, name)
	if err != nil {
		return err
	}
	dataResponse, err := s.Client.UpdateBinaryData(formData, dataID)
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
	tokenResponse, err := utils.LoadUserDataFromDosk()
	if err != nil {
		return err
	}
	s.Client.SetToken(tokenResponse.Token)
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
