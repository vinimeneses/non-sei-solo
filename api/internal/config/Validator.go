package config

import "github.com/go-playground/validator/v10"

// CustomValidator is a wrapper for validator.Validate
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate implements echo.Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
