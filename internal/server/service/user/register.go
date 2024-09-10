package user

import (
	"context"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

func (s *UserService) RegisterUser(ctx context.Context, user model.RegisterUser) (model.User, error) {
	_, err := s.Storage.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return model.User{}, utils.ErrEmailAlreadyUsed
	}
	passwordHash, err := utils.GetPasswordHash(user.Password, s.Config.SecretKey)
	if err != nil {
		return model.User{}, err
	}
	newUser, err := s.Storage.RegisterUser(ctx, user.Email, passwordHash)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}
