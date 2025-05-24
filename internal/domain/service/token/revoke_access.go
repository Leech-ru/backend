package token

import (
	"context"
	"github.com/google/uuid"
)

// RevokeAccessToken revoke access token.
func (s *tokenService) RevokeAccessToken(ctx context.Context, userID uuid.UUID) error {
	err := s.accessTokenRepo.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
