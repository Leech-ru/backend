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

	users, totalItems, err := s.userRepo.GetAllByFilter(
		ctx,
		limit,
		offset,
		req.Role,
		req.Query,
		req.EmailPrefix,
	)
	if err != nil {
		return nil, err
	}

	respItems := make([]*dto.User, 0, len(users))
	for _, user := range users {
		respItems = append(respItems, &dto.User{
			ID:      user.ID,
			Email:   user.Email,
			Name:    user.Name,
			Surname: user.Surname,
			Role:    user.Role,
		})
	}

	totalPages := 0
	if totalItems > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}
	currentPage := (offset / limit) + 1

	resp := dto.GetAllByFilterUsersResponse{
		Items: respItems,
		Pagination: dto.PaginationInfo{
			TotalItems:  totalItems,
			TotalPages:  totalPages,
			CurrentPage: currentPage,
			HasNext:     currentPage < totalPages,
			HasPrevious: currentPage > 1,
		},
	}

	return &resp, nil
}
