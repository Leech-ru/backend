package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// GetByID returns the category by ID.
func (s *categoryService) GetByID(ctx context.Context, req *dto.GetByIdCategoryRequest) (*dto.GetByIdCategoryResponse, error) {
	category, err := s.categoryRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CategoryNotFound):
		return nil, errorz.CategoryNotFound
	case err != nil:
		return nil, err
	}
	return &dto.GetByIdCategoryResponse{
		ID:      category.ID,
		ImageID: category.ImageID,
		Name:    category.Name,
	}, nil
}
