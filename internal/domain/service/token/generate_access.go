package token

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// GenerateAccessToken generate access token with userID
func (s *tokenService) GenerateAccessToken(ctx context.Context, userID uuid.UUID) (string, error) {
	token, jti, err := s.jwtService.GenerateToken(userID, s.jwtConfig.AccessTokenExpires())
	if err != nil {
		return "", err
	}

	err = s.accessTokenRepo.Set(ctx, userID, jti, time.Now().Add(s.jwtConfig.AccessTokenExpires()))
	if err != nil {
		return "", err
	}

	return token, nil
}
