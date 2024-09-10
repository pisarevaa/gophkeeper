package model

import (
	"time"

	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

type RegisterUser struct {
	Email    string `json:"email"    validate:"required,email,lte=250"`
	Password string `json:"password" validate:"required,lte=130"`
}

type UserResponse struct {
	ID        int64          `json:"id"`
	Email     string         `json:"email"`
	CreatedAt utils.Datetime `json:"createdAt"`
}

type User struct {
	ID        int64
	Email     string
	Password  string
	CreatedAt time.Time
}

type Login struct {
	Email    string `json:"email"    validate:"required,email,lte=250"`
	Password string `json:"password" validate:"required,lte=130"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
