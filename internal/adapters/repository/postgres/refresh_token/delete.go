package refresh_token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/refreshtoken"
	"Leech-ru/pkg/ent/user"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// DeleteByUserID removes a token by user ID
func (s *tokenRepo) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	_, err := s.client.RefreshToken.
		Delete().
		Where(refreshtoken.HasUserWith(user.IDEQ(userID))).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.TokenNotFound
	case err != nil:
		return fmt.Errorf("failed to query db: %w", err)
	}

	return nil
}
