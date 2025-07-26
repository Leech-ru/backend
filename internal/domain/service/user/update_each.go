package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/types"
	"Leech-ru/internal/domain/utils/password"
	"context"
	"errors"
)

// UpdateEach updates current user by ID.
//
// A user cannot modify another user who has the same or higher role.
//
// A user cannot assign a role that is equal to or higher than their own.
//
// SuperAdmin has no restrictions.
func (s *userService) UpdateEach(ctx context.Context, req *dto.UpdateEachUserRequest) (*dto.UpdateEachUserResponse, error) {
	userToUpdate, err := s.userRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.UserNotFound):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}

	if req.RequesterRole <= userToUpdate.Role && req.RequesterRole != types.RoleSuperAdmin {
		return nil, errorz.PermissionDenied
	}
	if req.Role != nil && req.RequesterRole != types.RoleSuperAdmin {
		if req.RequesterRole <= *req.Role {
			return nil, errorz.PermissionDenied
		}
	}

	if req.Name != nil {
		userToUpdate.Name = *req.Name
	}
	if req.Surname != nil {
		userToUpdate.Surname = *req.Surname
	}
	if req.Email != nil {
		userToUpdate.Email = *req.Email
	}
	if req.Password != nil {
		passwordHash, err := password.PasswordHash(*req.Password)
		if err != nil {
			return nil, err
		}
		userToUpdate.Password = passwordHash
	}
	if req.Role != nil {
		userToUpdate.Role = *req.Role
	}

	updatedUser, err := s.userRepo.Update(ctx, *userToUpdate)
	switch {
	case errors.Is(err, errorz.UserNotFound):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}

	return &dto.UpdateEachUserResponse{
		ID:      updatedUser.ID,
		Email:   updatedUser.Email,
		Name:    updatedUser.Name,
		Surname: updatedUser.Surname,
		Role:    updatedUser.Role,
	}, nil
}
