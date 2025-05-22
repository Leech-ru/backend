package validator

import (
	"github.com/go-playground/validator/v10"
)

func validateRole(fl validator.FieldLevel) bool {
	role := fl.Field().Int()
	switch role {
	case 0, 1, 2:
		return true
	default:
		return false
	}
}
