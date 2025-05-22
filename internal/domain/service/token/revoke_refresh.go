package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"errors"
	"github.com/google/uuid"
)

// RevokeRefreshToken revoke refresh token.
func (s *tokenService) RevokeRefreshToken(ctx context.Context, token string) (uuid.UUID, error) {
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

	tokenEntity.Jti = "-"
	_, err = s.refreshTokenRepo.Update(ctx, *tokenEntity)
	switch {
	case errors.Is(err, errorz.TokenNotFound):
		return uuid.Nil, errorz.Unauthorized
	case err != nil:
		return uuid.Nil, err
	}
	return userID, nil
}
