package validator

import (
	"github.com/go-playground/validator/v10"
)

func validatePackage(fl validator.FieldLevel) bool {
	packag := fl.Field().Int()
	switch packag {
	case 0, 1, 2:
		return true
	default:
		return false
	}
}
