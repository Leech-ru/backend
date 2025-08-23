package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
)

// Update updates an existing user
func (s *userRepo) Update(ctx context.Context, userEntity ent.User) (*ent.User, error) {
	updated, err := s.client.User.
		UpdateOneID(userEntity.ID).
		SetEmail(userEntity.Email).
		SetPassword(userEntity.Password).
		SetName(userEntity.Name).
		SetSurname(userEntity.Surname).
		SetRole(userEntity.Role).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.UserNotFound
	case ent.IsConstraintError(err):
		return nil, errorz.EmailAlreadyExist
	case err != nil:
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	return updated, nil
}
