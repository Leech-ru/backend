package user

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// Get returns list of users.
func (s *userService) GetAll(ctx context.Context, req *dto.GetAllUsersRequest) ([]*dto.GetUserResponse, error) {
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}

	users, err := s.userRepo.GetAll(ctx, limit, offset)
	switch {
	case err != nil:
		return nil, err
	}

	var userList []*dto.GetUserResponse
	for _, user := range users {
		userList = append(userList, &dto.GetUserResponse{
			ID:      user.ID,
			Email:   user.Email,
			Name:    user.Name,
			Surname: user.Surname,
			Role:    user.Role,
		})
	}
	return userList, err
}
