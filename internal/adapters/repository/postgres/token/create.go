package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new user in the database
func (s *tokenRepo) Create(ctx context.Context, entity ent.Token) (*ent.Token, error) {
	created, err := s.client.Token.
		Create().
		SetUserID(entity.UserID).
		SetToken(entity.Token).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.EmailAlreadyExist
	case err != nil:
		return nil, err
	}

	return created, nil
}
