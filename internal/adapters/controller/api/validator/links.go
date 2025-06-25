package validator

import (
	"github.com/go-playground/validator/v10"
	"net/url"
	"regexp"
	"strings"
)

var (
	ozonRegex        = regexp.MustCompile(`^(https?://)?(www\.)?ozon\.ru/product/.+`)
	wildberriesRegex = regexp.MustCompile(`^(https?://)?(www\.)?wildberries\.ru/catalog/\d+/detail\.aspx$`)
)

func validateOzonLink(fl validator.FieldLevel) bool {
	link := fl.Field().String()
	if link == "" {
		return true // обрабатывается через "omitempty"
	}
	parsed, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}
	host := parsed.Hostname()
	return strings.HasSuffix(host, "ozon.ru")
}

func validateWildberriesLink(fl validator.FieldLevel) bool {
	link := fl.Field().String()
	if link == "" {
		return true
	}
	parsed, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}
	host := parsed.Hostname()
	return strings.HasSuffix(host, "wildberries.ru")
}
