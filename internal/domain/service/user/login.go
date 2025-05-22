package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/utils/password"
	"context"
	"errors"
)

// Login returns the user by email.
func (s *userService) Login(ctx context.Context, req *dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	u, err := s.userRepo.GetByEmail(ctx, req.Email)
	switch {
	case errors.As(err, &errorz.UserNotFound):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}
	if !password.VerifyPassword(u.Password, req.Password) {
		return nil, errorz.PasswordMismatch
	}

	return &dto.LoginUserResponse{
		Token: "",
		User: dto.User{
			ID:      u.ID,
			Email:   u.Email,
			Name:    u.Name,
			Surname: u.Surname,
			Role:    u.Role,
		},
	}, nil
}
