package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Delete delete the user by ID.
func (s *userService) Delete(ctx context.Context, req *dto.DeleteUserRequest) error {
	err := s.userRepo.Delete(ctx, req.ID)
	switch {
	case errors.As(err, &errorz.UserNotFound):
		return errorz.UserNotFound
	case err != nil:
		return err
	}

	return nil
}
