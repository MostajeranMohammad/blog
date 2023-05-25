package utils

import "github.com/go-playground/validator/v10"

func ValidateDto(d interface{}) error {
	validate := validator.New()
	return validate.Struct(d)
}
