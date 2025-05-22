package refresh_token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

//TODO чета не работает апдейт

// Update updates token by existing user.
func (s *tokenRepo) Update(ctx context.Context, entity ent.RefreshToken) (*ent.RefreshToken, error) {
	updated, err := s.client.RefreshToken.
		UpdateOneID(entity.ID).
		SetJti(entity.Jti).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.TokenNotFound
	case err != nil:
		return nil, err
	}

	return updated, nil
}
