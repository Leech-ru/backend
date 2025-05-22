package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"errors"
	"github.com/google/uuid"
)

// ParseRefreshToken parse refresh token and returns user uuid or err if invalid..
func (s *tokenService) ParseRefreshToken(ctx context.Context, token string) (uuid.UUID, error) {
	userID, jti, err := s.jwtService.ParseToken(token)
	if err != nil {
		return uuid.Nil, errorz.InvalidToken
	}

	tokenEntity, err := s.refreshTokenRepo.GetByUserID(ctx, userID)
	switch {
	case errors.Is(err, errorz.TokenNotFound):
		return uuid.Nil, errorz.Unauthorized
	case err != nil:
		return uuid.Nil, err
	}
	if tokenEntity.Jti != jti {
		return uuid.Nil, errorz.Unauthorized
	}

	return userID, nil
}
