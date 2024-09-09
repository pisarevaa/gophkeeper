package model

import (
	"time"
)

type RegisterUser struct {
	Email    string `json:"email"    validate:"required,email,lte=250"`
	Password string `json:"password" validate:"required,lte=130"`
}

type User struct {
	ID        int64
	Email     string
	Password  string
	CreatedAt time.Time
}
