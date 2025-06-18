package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Logout delete all sessions
func (s *userService) Logout(ctx context.Context, req *dto.LogoutRequest) error {
	err := s.tokenService.LogoutAllSessions(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.Unauthorized):
		return errorz.Unauthorized
	case err != nil:
		return err
	}
	return nil
}
