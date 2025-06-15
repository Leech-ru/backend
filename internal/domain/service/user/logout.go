package user

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// Logout delete all sessions
func (s *userService) Logout(ctx context.Context, req *dto.LogoutRequest) error {
	err := s.tokenService.LogoutAllSessions(ctx, req.ID)
	return err
}
