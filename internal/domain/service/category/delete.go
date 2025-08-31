package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Delete delete category by ID.
func (s *categoryService) Delete(ctx context.Context, req *dto.DeleteCategoryRequest) error {
	category, err := s.categoryRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CategoryNotFound):
		return errorz.CategoryNotFound
	case err != nil:
		return err
	}

	if err := s.imageService.Delete(ctx, &dto.DeleteImageRequest{ID: category.ImageID}); err != nil {
		return err
	}

	err = s.categoryRepo.Delete(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CategoryNotFound):
		return errorz.CategoryNotFound
	case err != nil:
		return err
	}

	return nil
}
