package service

import (
	"github.com/pisarevaa/gophkeeper/internal/agent/utils"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Регистрация нового пользователя.
func (s *Service) RegisterUser(user model.RegisterUser) error {
	if err := s.Validator.Struct(user); err != nil {
		return err
	}
	_, err := s.Client.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}

// Авторизация пользователя.
func (s *Service) LoginUser(user model.RegisterUser) error {
	if err := s.Validator.Struct(user); err != nil {
		return err
	}
	tokenResponse, err := s.Client.LoginUser(user)
	if err != nil {
		return err
	}
	if errSave := utils.SaveUserDataToDisk(tokenResponse); errSave != nil {
		return err
	}
	return nil
}
