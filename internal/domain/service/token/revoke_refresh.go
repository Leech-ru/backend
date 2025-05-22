package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"github.com/google/uuid"
)

// RevokeRefreshToken revoke refresh token.
func (s *tokenService) RevokeRefreshToken(ctx context.Context, token string) (uuid.UUID, error) {
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
	tokenEntity.Jti = "-"
	_, err = s.refreshTokenRepo.Update(ctx, *tokenEntity)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, err
}
