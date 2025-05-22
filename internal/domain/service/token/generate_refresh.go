package token

import (
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

// GenerateRefreshToken generate refresh token with userID
func (s *tokenService) GenerateRefreshToken(ctx context.Context, userID uuid.UUID) (string, error) {
	token, jti, err := s.jwtService.GenerateToken(userID, s.jwtConfig.RefreshTokenExpires())
	if err != nil {
		return "", err
	}

	_, err = s.refreshTokenRepo.Upsert(ctx, ent.RefreshToken{
		Jti: jti,
		Edges: ent.RefreshTokenEdges{
			User: &ent.User{ID: userID},
		},
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
