package token

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

func (s *tokenService) LogoutAllSessions(ctx context.Context, userID uuid.UUID) error {
	tokenEntity, err := s.refreshTokenRepo.GetByUserID(ctx, userID)
	switch {
	case errors.Is(err, errorz.TokenNotFound):
		return nil
	case err != nil:
		return err
	}

	tokenEntity.Jti = "-"
	_, err = s.refreshTokenRepo.Update(ctx, *tokenEntity)
	switch {
	case errors.Is(err, errorz.TokenNotFound):
		return nil
	case err != nil:
		return err
	}

	err = s.accessTokenRepo.Set(ctx, userID, "-", time.Now().Add(s.jwtConfig.AccessTokenExpires()))
	if err != nil {
		return err
	}

	return nil
}
