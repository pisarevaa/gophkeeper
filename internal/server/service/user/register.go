package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

func (s *UserService) RegisterUser(ctx context.Context, user model.RegisterUser) (model.User, int, error) {
	_, err := s.Storage.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return model.User{}, http.StatusConflict, errors.New("email уже использован")
	}
	passwordHash, err := utils.GetPasswordHash(user.Password, s.Config.Security.SecretKey)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	newUser, err := s.Storage.RegisterUser(ctx, user.Email, passwordHash)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return newUser, 0, nil
}

func (s *UserService) Login(ctx context.Context, user model.RegisterUser) (string, int, error) {
	foundUser, err := s.Storage.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", http.StatusNotFound, errors.New("email не найден")
	}
	isCorrect, err := utils.CheckPasswordHash(user.Password, foundUser.Password, s.Config.Security.SecretKey)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	if !isCorrect {
		return "", http.StatusUnauthorized, errors.New("некорректный пароль")
	}

	token, err := utils.GenerateJWTString(s.Config.Security.TokenExpSec, s.Config.Security.SecretKey, foundUser.Email)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return token, 0, nil
}
