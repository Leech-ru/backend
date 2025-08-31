package category

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll realizes a search for cosmetics with filtering parameters.
func (s *categoryService) GetAll(ctx context.Context, req *dto.GetAllCategoriesRequest) (*dto.GetAllCategoriesResponse, error) {
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}
	allCategories, err := s.categoryRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	var resp dto.GetAllCategoriesResponse
	for _, category := range allCategories {
		resp = append(resp, &dto.Category{
			ID:      category.ID,
			ImageID: category.ImageID,
			Name:    category.Name,
		})
	}
	return &resp, nil
}
