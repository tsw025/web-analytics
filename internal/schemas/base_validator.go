package schemas

import "github.com/go-playground/validator/v10"

type BaseValidator struct {
	Validator *validator.Validate
}

func (bv *BaseValidator) Validate(i interface{}) error {
	return bv.Validator.Struct(i)
}
