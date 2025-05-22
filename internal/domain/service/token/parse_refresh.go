package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"github.com/google/uuid"
)

// ParseRefreshToken parse refresh token and returns user uuid or err if invalid..
func (s *tokenService) ParseRefreshToken(ctx context.Context, token string) (uuid.UUID, error) {
	userID, jti, err := s.jwtService.ParseToken(token)
	if err != nil {
		return uuid.Nil, err
	}

	tokenEntity, err := s.refreshTokenRepo.GetByUserID(ctx, userID)
	if err != nil {
		return uuid.Nil, err
	}
	if tokenEntity.Jti != jti {
		return uuid.Nil, errorz.Unauthorized
	}

	return userID, err
}
