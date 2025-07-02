package user

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAllByFilter Realizes the search for users with filtering parameters.
func (s *userService) GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterUsersRequest) (*dto.GetAllByFilterUsersResponse, error) {
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}

	users, err := s.userRepo.GetAllByFilter(ctx, limit, offset, req.Role, req.NamePrefix, req.SurnamePrefix, req.EmailPrefix)
	if err != nil {
		return nil, err
	}

	var resp dto.GetAllByFilterUsersResponse
	for _, user := range users {
		resp = append(resp, &dto.User{
			ID:      user.ID,
			Email:   user.Email,
			Name:    user.Name,
			Surname: user.Surname,
			Role:    user.Role,
		})
	}

	return &resp, nil
}
