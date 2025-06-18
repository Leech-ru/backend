package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/types"
	"Leech-ru/internal/domain/utils/password"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
)

// Register returns registered user with token.
func (s *userService) Register(ctx context.Context, req *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	passwordHash, err := password.PasswordHash(req.Password)
	if err != nil {
		return nil, err
	}
	user := ent.User{
		Email:    req.Email,
		Password: passwordHash,
		Name:     req.Name,
		Surname:  req.Surname,
		Role:     types.RoleUser,
	}
	if s.serverConfig.DevMode() && req.Role != nil {
		user.Role = *req.Role
	}
	u, err := s.userRepo.Create(ctx, user)
	switch {
	case errors.Is(err, errorz.EmailAlreadyExist):
		return nil, errorz.EmailAlreadyExist
	case err != nil:
		return nil, err
	}

	token, err := s.tokenService.GenerateRefreshToken(ctx, u.ID)
	if err != nil {
		return nil, err
	}
	err = s.tokenService.RevokeAccessToken(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponse{
		RefreshToken: token,
		User: dto.User{
			ID:      u.ID,
			Email:   u.Email,
			Name:    u.Name,
			Surname: u.Surname,
			Role:    u.Role,
		},
	}, nil
}
