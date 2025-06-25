package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"reflect"
	"strings"
)

type Validator struct {
	validator *validator.Validate
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func New() *Validator {
	newValidator := validator.New()

	_ = newValidator.RegisterValidation("minleechsum", validateMinLeechSum)
	_ = newValidator.RegisterValidation("maxleechsum", validateMaxLeechSum)

	_ = newValidator.RegisterValidation("role", validateRole)

	_ = newValidator.RegisterValidation("category", validateCategory)

	_ = newValidator.RegisterValidation("ozonlink", validateOzonLink)
	_ = newValidator.RegisterValidation("wildberrieslink", validateWildberriesLink)

	return &Validator{
		newValidator,
	}
}

func (v Validator) ValidateData(data interface{}) *echo.HTTPError {
	var validationErrors []ErrorResponse

	if data == nil {
		return &echo.HTTPError{
			Code:    echo.ErrBadRequest.Code,
			Message: "validation failed: input data is nil",
		}
	}

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			if err := v.validator.Struct(val.Index(i).Interface()); err != nil {
				var invalidValidationError *validator.InvalidValidationError
				if errors.As(err, &invalidValidationError) {
					return &echo.HTTPError{
						Code:    echo.ErrBadRequest.Code,
						Message: err.Error(),
					}
				}

				var validationErrs validator.ValidationErrors
				ok := errors.As(err, &validationErrs)
				if !ok {
					return &echo.HTTPError{
						Code:    echo.ErrBadRequest.Code,
						Message: "unexpected validation error type",
					}
				}

				for _, err := range validationErrs {
					var elem ErrorResponse
					elem.FailedField = fmt.Sprintf("[%d].%s", i, err.Field())
					elem.Tag = err.Tag()
					elem.Value = err.Value()
					elem.Error = true
					validationErrors = append(validationErrors, elem)
				}
			}
		}
	} else {
		if err := v.validator.Struct(data); err != nil {
			var invalidValidationError *validator.InvalidValidationError
			if errors.As(err, &invalidValidationError) {
				return &echo.HTTPError{
					Code:    echo.ErrBadRequest.Code,
					Message: err.Error(),
				}
			}

			var validationErrs validator.ValidationErrors
			ok := errors.As(err, &validationErrs)
			if !ok {
				return &echo.HTTPError{
					Code:    echo.ErrBadRequest.Code,
					Message: "unexpected validation error type",
				}
			}

			for _, err := range validationErrs {
				var elem ErrorResponse
				elem.FailedField = err.Field()
				elem.Tag = err.Tag()
				elem.Value = err.Value()
				elem.Error = true
				validationErrors = append(validationErrors, elem)
			}
		}
	}

	if len(validationErrors) > 0 && validationErrors[0].Error {
		errMessages := make([]string, 0)
		for _, err := range validationErrors {
			errMessages = append(errMessages, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return &echo.HTTPError{
			Code:    echo.ErrBadRequest.Code,
			Message: strings.Join(errMessages, " and "),
		}
	}
	return nil
}
