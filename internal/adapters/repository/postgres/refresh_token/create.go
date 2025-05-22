package refresh_token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new user in the database
func (s *tokenRepo) Create(ctx context.Context, entity ent.RefreshToken) (*ent.RefreshToken, error) {
	created, err := s.client.RefreshToken.
		Create().
		SetJti(entity.Jti).
		SetUserID(entity.Edges.User.ID).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.UserAlreadyHasToken
	case err != nil:
		return nil, err
	}

	return created, nil
}
