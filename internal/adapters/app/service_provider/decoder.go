package service_provider

import (
	"github.com/go-playground/form"
)

func (s *ServiceProvider) Decoder() *form.Decoder {
	if s.decoder == nil {
		s.decoder = form.NewDecoder()
	}

	return s.decoder
}
