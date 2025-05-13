package token

import (
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

// NewToken create token for user id.
func (s *tokenService) NewToken(ctx context.Context, id uuid.UUID) (string, error) {
	tokenString, err := s.jwtService.GenerateToken(id)
	if err != nil {
		return "", err
	}

	existingToken, err := s.tokenRepo.Create(ctx, ent.Token{
		UserID: id,
		Token:  tokenString,
	})
	if err != nil {
		return "", err
	}

	return existingToken.Token, nil
}
