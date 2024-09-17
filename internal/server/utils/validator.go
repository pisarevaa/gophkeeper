package utils

import (
	"github.com/go-playground/validator/v10"
)

// Валидатор данных структур.
func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}
