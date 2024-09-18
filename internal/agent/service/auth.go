package service

import (
	"errors"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/agent/model"
)

// Регистрация нового пользователя.
func (s *Service) RegisterUser(user model.RegisterUser) error {
	if err := s.Validator.Struct(user); err != nil {
		return err
	}
	status, err := s.Client.RegisterUser(user)
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
