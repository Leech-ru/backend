package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Delete removes a user by ID
func (s *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.client.User.
		DeleteOneID(id).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.UserNotFound
	case err != nil:
		return fmt.Errorf("failed to query db: %w", err)
	}

	return nil
}
