package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/token"
	"context"
	"github.com/google/uuid"
)

// GetById retrieves a user token by ID
func (s *tokenRepo) GetById(ctx context.Context, id uuid.UUID) (*ent.Token, error) {
	u, err := s.client.Token.
		Query().
		Where(token.UserID(id)).
		Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, errorz.TokenNotFound
	case err != nil:
		return nil, err
	}

	return u, nil
}

// GetByToken retrieves a user ID by token
func (s *tokenRepo) GetByToken(ctx context.Context, inputToken string) (*ent.Token, error) {
	t, err := s.client.Token.
		Query().
		Where(token.Token(inputToken)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.TokenNotFound
	case err != nil:
		return nil, err
	}

	return t, nil
}
