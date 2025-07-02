package validator

import "github.com/go-playground/validator/v10"

func validateWeekday(fl validator.FieldLevel) bool {
	weekday := fl.Field().Int()
	switch weekday {
	case 0, 1, 2, 3, 4, 5, 6:
		return true
	default:
		return false
	}
}
