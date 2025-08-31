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

	if req.ImageID != nil {
		if cond, err := s.imageService.Exists(ctx, *req.ImageID); err != nil || !cond {
			switch {
			case !cond:
				return nil, errorz.ImageNotFound
			default:
				return nil, err
			}
		}
		if err := s.imageService.Delete(ctx, &dto.DeleteImageRequest{ID: categoryToUpdate.ImageID}); err != nil {
			switch {
			case errors.Is(err, errorz.ImageNotFound):
				return nil, errorz.ImageNotFound
			case err != nil:
				return nil, err
			}
		}
		categoryToUpdate.ImageID = *req.ImageID
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
	case errors.Is(err, errorz.ImageNotFound):
		return nil, errorz.ImageNotFound
	case err != nil:
		return nil, err
	}

	return &dto.UpdateCategoryResponse{
		ID:      updatedCategory.ID,
		ImageID: updatedCategory.ImageID,
		Name:    updatedCategory.Name,
	}, nil
}
