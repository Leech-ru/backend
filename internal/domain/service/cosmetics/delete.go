package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Delete removes cosmetics and its image (if any)
func (s *cosmeticsService) Delete(ctx context.Context, req *dto.DeleteCosmeticsRequest) error {
	cosmetics, err := s.cosmeticsRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return errorz.CosmeticsNotFound
	case err != nil:
		return err
	}

	if err := s.imageService.Delete(ctx, &dto.DeleteImageRequest{ID: cosmetics.ImageID}); err != nil {
		return err
	}

	if err := s.cosmeticsRepo.Delete(ctx, req.ID); err != nil {
		return err
	}

	return nil
}
