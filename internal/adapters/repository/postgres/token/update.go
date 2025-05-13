package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Update updates an existing user
func (s *tokenRepo) Update(ctx context.Context, tokenEntity ent.Token) (*ent.Token, error) {
	updated, err := s.client.Token.
		UpdateOneID(tokenEntity.ID).
		SetToken(tokenEntity.Token).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.TokenNotFound
	case err != nil:
		return nil, err
	}

	return updated, nil
}
