package validator

import (
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

// validateMaxFileSize checks that file size <= max (bytes)
// Usage: validate:"maxfilesize=5242880"  // 5 MB
func validateMaxFileSize(fl validator.FieldLevel) bool {
	size, ok := fl.Field().Interface().(int64)
	if !ok {
		return false
	}

	limit, err := strconv.ParseInt(fl.Param(), 10, 64)
	if err != nil {
		return false
	}

	return 0 <= size && size <= limit
}

// validateFileType checks that ContentType matches allowed MIME list exactly, case-insensitive
// Usage: validate:"filetype=image/png;image/jpeg;image/gif"
func validateFileType(fl validator.FieldLevel) bool {
	contentType, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	contentType = strings.ToLower(strings.TrimSpace(contentType))

	allowed := strings.Split(fl.Param(), ";")
	for _, t := range allowed {
		if contentType == strings.ToLower(strings.TrimSpace(t)) {
			return true
		}
	}
	return false
}
