package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
)

// Create category and returns it.
func (s *categoryService) Create(ctx context.Context, req *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, error) {
	category := &ent.Category{
		Name: req.Name,
	}

	category, err := s.categoryRepo.Create(ctx, *category)
	switch {
	case errors.Is(err, errorz.InvalidCategoryFormat):
		return nil, errorz.InvalidCategoryFormat
	case err != nil:
		return nil, err
	}
	return &dto.CreateCategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}
