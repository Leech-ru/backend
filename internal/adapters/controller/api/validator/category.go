package validator

import "github.com/go-playground/validator/v10"

func validateCategory(fl validator.FieldLevel) bool {
	category := fl.Field().Int()
	switch category {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
		return true
	default:
		return false
	}
}
