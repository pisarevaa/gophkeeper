package model

type RegisterUser struct {
	Email    string `json:"email"    validate:"required,email,lte=250"`
	Password string `json:"password"                                   binding:"required,lte=130"`
}
