package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"errors"
	"github.com/google/uuid"
)

// GetIDByToken get user id by token.
func (s *tokenService) GetIDByToken(ctx context.Context, token string) (uuid.UUID, error) {
	existingToken, err := s.tokenRepo.GetByToken(ctx, token)
	switch {
	case errors.As(err, &errorz.TokenNotFound):
		return uuid.Nil, errorz.TokenNotFound
	case err != nil:
		return uuid.Nil, err
	}

	return existingToken.UserID, nil
}
