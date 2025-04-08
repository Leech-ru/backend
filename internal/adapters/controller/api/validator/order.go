package validator

import (
	"LutiLeech/internal/domain/dto"
	"github.com/go-playground/validator/v10"
	"strconv"
)

func validateMinLeechSum(fl validator.FieldLevel) bool {
	orderDetails, ok := fl.Field().Interface().(dto.OrderDetails)
	if !ok {
		return false
	}
	sum := getLeechSum(orderDetails)
	minSum, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}
	return sum >= minSum
}

func validateMaxLeechSum(fl validator.FieldLevel) bool {
	orderDetails, ok := fl.Field().Interface().(dto.OrderDetails)
	if !ok {
		return false
	}

	sum := getLeechSum(orderDetails)
	maxSum, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}

	return sum <= maxSum
}

func getLeechSum(od dto.OrderDetails) int {
	sum := 0
	if od.LeechSize1 != nil {
		sum += *od.LeechSize1
	}
	if od.LeechSize2 != nil {
		sum += *od.LeechSize2
	}
	if od.LeechSize3 != nil {
		sum += *od.LeechSize3
	}
	return sum
}
