package refresh_token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/refreshtoken"
	"Leech-ru/pkg/ent/user"
	"context"
	"github.com/google/uuid"
)

// GetByUserID retrieves a user token by ID
func (s *tokenRepo) GetByUserID(ctx context.Context, userID uuid.UUID) (*ent.RefreshToken, error) {
	token, err := s.client.RefreshToken.
		Query().
		Where(refreshtoken.HasUserWith(user.IDEQ(userID))).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.TokenNotFound
	case err != nil:
		return nil, err
	}

	return token, nil
}
