package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"errors"
	"github.com/google/uuid"
)

// ParseAccessToken parse access token and returns user uuid or err if invalid.
func (s *tokenService) ParseAccessToken(ctx context.Context, token string) (uuid.UUID, error) {
	userID, jti, err := s.jwtService.ParseToken(token)
	if err != nil {
		return uuid.Nil, errorz.InvalidToken
	}

	savedJti, err := s.accessTokenRepo.Get(ctx, userID)
	switch {
	case errors.Is(err, errorz.TokenNotFound):
		return uuid.Nil, errorz.Unauthorized
	case err != nil:
		return uuid.Nil, err
	}
	if savedJti != jti {
		return uuid.Nil, errorz.Unauthorized
	}

	return userID, nil
}
