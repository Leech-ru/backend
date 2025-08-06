package category

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Delete delete category by ID.
func (s *categoryService) Delete(ctx context.Context, req *dto.DeleteCategoryRequest) error {
	err := s.categoryRepo.Delete(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CategoryNotFound):
		return errorz.CategoryNotFound
	case err != nil:
		return err
	}

	return nil
}
