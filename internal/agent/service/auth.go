package service

import (
	"errors"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/agent/utils"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Регистрация нового пользователя.
func (s *Service) RegisterUser(user model.RegisterUser) error {
	if err := s.Validator.Struct(user); err != nil {
		return err
	}
	_, status, err := s.Client.RegisterUser(user)
	if err != nil {
		return err
	}
	switch status {
	case http.StatusOK:
		return nil
	case http.StatusUnprocessableEntity:
		return errors.New("incorrect data provided")
	case http.StatusConflict:
		return errors.New("email is already used")
	default:
		return errors.New("internal server error")
	}
}

// Авторизация пользователя.
func (s *Service) LoginUser(user model.RegisterUser) error {
	if err := s.Validator.Struct(user); err != nil {
		return err
	}
	tokenResponse, status, err := s.Client.LoginUser(user)
	if err != nil {
		return err
	}
	switch status {
	case http.StatusOK:
		if errSave := utils.SaveUserDataToDisk(tokenResponse); errSave != nil {
			return err
		}
		return nil
	case http.StatusUnprocessableEntity:
		return errors.New("incorrect data provided")
	case http.StatusConflict:
		return errors.New("email is already used")
	default:
		return errors.New("internal server error")
	}
}
