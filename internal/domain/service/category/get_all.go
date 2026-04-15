package category

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll realizes a search for cosmetics with filtering parameters.
func (s *categoryService) GetAll(ctx context.Context, req *dto.GetAllCategoriesRequest) (*dto.GetAllCategoriesResponse, error) {
	limit := 10
	if req.Limit != nil && *req.Limit > 0 {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}
	allCategories, totalItems, err := s.categoryRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	respItems := make([]*dto.Category, 0, len(allCategories))
	for _, category := range allCategories {
		respItems = append(respItems, &dto.Category{
			ID:      category.ID,
			ImageID: category.ImageID,
			Name:    category.Name,
		})
	}

	totalPages := 0
	if totalItems > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}
	currentPage := (offset / limit) + 1

	resp := dto.GetAllCategoriesResponse{
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
