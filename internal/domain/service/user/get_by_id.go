package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// GetByID returns the user by ID.
func (s *userService) GetByID(ctx context.Context, req *dto.GetUserRequest) (*dto.GetUserResponse, error) {
	u, err := s.userRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.UserNotFound):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}

	return &dto.GetUserResponse{
		ID:      u.ID,
		Email:   u.Email,
		Name:    u.Name,
		Surname: u.Surname,
		Role:    u.Role,
	}, nil
}
