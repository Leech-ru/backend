package service_provider

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/service/user"
	"context"
)

type userService interface {
	Register(ctx context.Context, req *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error)
	Login(ctx context.Context, req *dto.LoginUserRequest) (*dto.LoginUserResponse, error)
	ChangePassword(ctx context.Context, req *dto.ChangePasswordRequest) (*dto.ChangePasswordResponse, error)
	Get(ctx context.Context, req *dto.GetUserRequest) (*dto.GetUserResponse, error)
	GetAll(ctx context.Context, req *dto.GetAllUsersRequest) (*dto.GetAllUsersResponse, error)
	Update(ctx context.Context, req *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	Delete(ctx context.Context, req *dto.DeleteUserRequest) error
}

func (s *ServiceProvider) UserService() userService {
	if s.userService == nil {
		s.userService = user.NewUserService(s.DB(), s.TokenService())
	}
	return s.userService
}
