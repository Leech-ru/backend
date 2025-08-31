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
		Name:    req.Name,
		ImageID: req.ImageID,
	}

	if cond, err := s.imageService.Exists(ctx, req.ImageID); err != nil || !cond {
		switch {
		case !cond:
			return nil, errorz.ImageNotFound
		case err != nil:
			return nil, err
		}
	}

	category, err := s.categoryRepo.Create(ctx, *category)
	switch {
	case errors.Is(err, errorz.InvalidCategoryFormat):
		return nil, errorz.InvalidCategoryFormat
	case err != nil:
		return nil, err
	}
	return &dto.CreateCategoryResponse{
		ID:      category.ID,
		Name:    category.Name,
		ImageID: category.ImageID,
	}, nil
}
