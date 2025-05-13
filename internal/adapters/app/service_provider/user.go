package service_provider

import (
	"Leech-ru/internal/domain/service/user"
)

type userService interface {
}

func (s *ServiceProvider) UserService() userService {
	if s.userService == nil {
		s.userService = user.NewUserService(s.DB(), s.TokenService())
	}
	return s.userService
}
