package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/types"
	"context"
	"errors"
	"github.com/google/uuid"
)

// GetRoleByID returns the user role by ID.
func (s *userService) GetRoleByID(ctx context.Context, userID uuid.UUID) (types.Role, error) {
	u, err := s.userRepo.GetById(ctx, userID)
	switch {
	case errors.As(err, &errorz.UserNotFound):
		return 0, errorz.UserNotFound
	case err != nil:
		return 0, err
	}
	return types.Role(u.Role), nil
}
