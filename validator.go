package main

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func setupValidator(v *validator.Validate) *validator.Validate {
	_ = v.RegisterValidation("datetime", validateDateTime)
	return v
}

func validateDateTime(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	// Convert time.Time value to string
	strValue := value.Format(time.RFC3339)

	if strValue == "0001-01-01T00:00:00Z" {
		return false
	}

	_, err := time.Parse(time.RFC3339, strValue)
	return err == nil
}
