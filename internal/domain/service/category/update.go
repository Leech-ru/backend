package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Update change cosmetics data.
func (s *categoryService) Update(ctx context.Context, req *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, error) {
	categoryToUpdate, err := s.categoryRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CategoryNotFound):
		return nil, errorz.CategoryNotFound
	case err != nil:
		return nil, err
	}
	if req.Name != nil {
		categoryToUpdate.Name = *req.Name
	}

	updatedCategory, err := s.categoryRepo.Update(ctx, *categoryToUpdate)
	switch {
	case errors.Is(err, errorz.CategoryNotFound):
		return nil, errorz.CategoryNotFound
	case errors.Is(err, errorz.InvalidCategoryFormat):
		return nil, errorz.InvalidCategoryFormat
	case err != nil:
		return nil, err
	}

	return &dto.UpdateCategoryResponse{
		ID:   updatedCategory.ID,
		Name: updatedCategory.Name,
	}, nil
}
