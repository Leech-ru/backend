package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/token"
	"context"
	"github.com/google/uuid"
)

// DeleteByUserID removes a user by ID
func (s *tokenRepo) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	_, err := s.client.Token.
		Delete().
		Where(token.UserID(userID)).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.TokenNotFound
	case err != nil:
		return err
	}

	return nil
}
