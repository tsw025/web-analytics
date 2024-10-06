package schemas

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type BaseValidator struct {
	Validator *validator.Validate
}

func (bv *BaseValidator) Validate(i interface{}) error {
	err := bv.Validator.RegisterValidation("password_format", passwordValidation)
	if err != nil {
		return errors.New("error registering password_format validation")
	}

	return bv.Validator.Struct(i)
}
