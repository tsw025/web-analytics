package schemas

import "github.com/go-playground/validator/v10"

type BaseValidator struct {
	Validator *validator.Validate
}

func (bv *BaseValidator) Validate(i interface{}) error {
	err := bv.Validator.RegisterValidation("password_val", passwordValidation)
	if err != nil {
		return err
	}

	return bv.Validator.Struct(i)
}
