package user

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll returns list of users.
func (s *userService) GetAll(ctx context.Context, req *dto.GetAllUsersRequest) (*dto.GetAllUsersResponse, error) {
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}

	users, err := s.userRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var userList dto.GetAllUsersResponse
	for _, user := range users {
		userList = append(userList, &dto.User{
			ID:      user.ID,
			Email:   user.Email,
			Name:    user.Name,
			Surname: user.Surname,
			Role:    user.Role,
		})
	}

	return &userList, nil
}
