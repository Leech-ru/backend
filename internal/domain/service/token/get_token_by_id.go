package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"errors"
	"github.com/google/uuid"
)

// GetTokenByID get token by user id.
func (s *tokenService) GetTokenByID(ctx context.Context, id uuid.UUID) (string, error) {
	existingToken, err := s.tokenRepo.GetById(ctx, id)
	switch {
	case errors.As(err, &errorz.TokenNotFound):
		return "", errorz.TokenNotFound
	case err != nil:
		return "", err
	}

	return existingToken.Token, nil
}
