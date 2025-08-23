package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/user"
	"context"
	"fmt"
)

// GetByEmail retrieves a user by email
func (s *userRepo) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := s.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	return u, nil
}
