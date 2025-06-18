package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/user"
	"context"
	"github.com/google/uuid"
)

// GetById retrieves a user by ID
func (s *userRepo) GetById(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	u, err := s.client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}

	return u, nil
}
