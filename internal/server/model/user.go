package model

import (
	"time"
)

type RegisterUser struct {
	Email    string `json:"email"    validate:"required,email,lte=250"`
	Password string `json:"password" validate:"required,gt=5,lte=250"`
}

type UserResponse struct {
	ID        int64    `json:"id"`
	Email     string   `json:"email"`
	CreatedAt DateTime `json:"createdAt"`
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
