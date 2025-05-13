package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
)

// TODO реализовать токен
func (s *userService) Register(ctx context.Context, req dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	user, err := s.userRepo.Create(ctx, ent.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
		Role:     req.Role,
	})
	switch {
	case errors.As(err, &errorz.EmailAlreadyExist):
		return nil, errorz.EmailAlreadyExist
	case err != nil:
		return nil, err
	}
	return &dto.RegisterUserResponse{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		Surname: user.Surname,
		Role:    user.Role,
	}, nil
}
