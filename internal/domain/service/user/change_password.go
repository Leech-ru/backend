package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/utils/password"
	"context"
	"errors"
)

// ChangePassword change the password by id.
func (s *userService) ChangePassword(ctx context.Context, req *dto.ChangePasswordRequest) (*dto.ChangePasswordResponse, error) {
	userToUpdate, err := s.userRepo.GetById(ctx, req.ID)
	switch {
	case errors.As(err, &errorz.UserNotFound):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}

	if req.OldPassword == req.NewPassword {
		return nil, errorz.PasswordsCoincidence
	}

	if !password.VerifyPassword(userToUpdate.Password, req.OldPassword) {
		return nil, errorz.PasswordMismatch
	}
	userToUpdate.Password, err = password.PasswordHash(req.NewPassword)
	if err != nil {
		return nil, err
	}
	_, err = s.userRepo.Update(ctx, *userToUpdate)
	switch {
	case errors.As(err, &errorz.UserNotFound):
		return nil, errorz.UserNotFound
	case err != nil:
		return nil, err
	}

	token, err := s.tokenService.UpdateToken(ctx, userToUpdate.ID)
	if err != nil {
		return nil, err
	}

	return &dto.ChangePasswordResponse{
		Token: token,
	}, nil
}
