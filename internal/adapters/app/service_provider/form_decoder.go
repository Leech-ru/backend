package service_provider

import (
	"github.com/go-playground/form"
)

func (s *ServiceProvider) Decoder() *form.Decoder {
	if s.formDecoder == nil {
		s.formDecoder = form.NewDecoder()
	}

	return s.formDecoder
}
