package model

// import (
// 	"time"
// )

type RegisterUser struct {
	Email    string `json:"email"    validate:"required,email,lte=250"`
	Password string `json:"password" validate:"required,gt=5,lte=250"`
}
