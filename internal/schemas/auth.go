package schemas

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,lowercase,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32,alphanum,lowercase"`
	Password string `json:"password" validate:"required,min=8,max=16,password_val"`
}

func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var hasMinLen = len(password) >= 8
	var hasMaxLen = len(password) <= 16
	var hasUpper = false
	var hasLower = false
	var hasNumber = false
	var hasSpecial = false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasMaxLen && hasUpper && hasLower && hasNumber && hasSpecial
}
