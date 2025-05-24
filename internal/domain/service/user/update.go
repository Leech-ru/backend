package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Update updates user by ID.
func (s *userService) Update(ctx context.Context, req *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	userToUpdate, err := s.userRepo.GetById(ctx, req.ID)
	switch {
	case errors.As(err, &errorz.UserNotFound):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}
	if req.Name != nil {
		userToUpdate.Name = *req.Name
	}
	if req.Surname != nil {
		userToUpdate.Surname = *req.Surname
	}

	updatedUser, err := s.userRepo.Update(ctx, *userToUpdate)
	switch {
	case errors.As(err, &errorz.UserNotFound):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}

	return &dto.UpdateUserResponse{
		ID:      updatedUser.ID,
		Email:   updatedUser.Email,
		Name:    updatedUser.Name,
		Surname: updatedUser.Surname,
		Role:    updatedUser.Role,
	}, nil
}
