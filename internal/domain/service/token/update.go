package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
	"github.com/google/uuid"
)

// UpdateToken set token by user id.
func (s *tokenService) UpdateToken(ctx context.Context, id uuid.UUID) (string, error) {
	tokenString, err := s.jwtService.GenerateToken(id)
	if err != nil {
		return "", err
	}
	tokenEntity := ent.Token{
		UserID: id,
		Token:  tokenString,
	}

	existingToken, err := s.tokenRepo.Update(ctx, tokenEntity)
	if errors.As(err, &errorz.TokenNotFound) {
		existingToken, err = s.tokenRepo.Create(ctx, tokenEntity)
		if err != nil {
			return "", err
		}
	} else {
		return "", err
	}

	return existingToken.Token, nil
}
