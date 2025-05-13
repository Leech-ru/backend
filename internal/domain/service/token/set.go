package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
	"github.com/google/uuid"
)

// SetToken set token by user id.
func (s *tokenService) SetToken(ctx context.Context, id uuid.UUID) (string, error) {
	tokenString, err := s.jwtService.GenerateToken(id)
	if err != nil {
		return "", err
	}
	existingToken, err := s.tokenRepo.GetById(ctx, id)
	if err != nil && !errors.As(err, &errorz.TokenNotFound) {
		return "", err
	}

	tokenEntity := ent.Token{
		UserID: id,
		Token:  tokenString,
	}

	if existingToken != nil {
		tokenEntity.ID = existingToken.ID
		existingToken, err = s.tokenRepo.Update(ctx, tokenEntity)
	} else {
		existingToken, err = s.tokenRepo.Create(ctx, tokenEntity)
	}

	if err != nil {
		return "", err
	}

	return existingToken.Token, nil
}
