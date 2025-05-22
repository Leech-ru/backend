package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

// RevokeAccessToken revoke access token.
func (s *tokenService) RevokeAccessToken(ctx context.Context, token string) (uuid.UUID, error) {
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

	err = s.accessTokenRepo.Set(ctx, userID, "-", time.Now().Add(s.jwtConfig.AccessTokenExpires()))
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}
