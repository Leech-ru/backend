package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Delete delete cosmetics by ID.
func (s *cosmeticsService) Delete(ctx context.Context, req *dto.DeleteCosmeticsRequest) error {
	err := s.cosmeticsRepo.Delete(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return errorz.CosmeticsNotFound
	case err != nil:
		return err
	}

	return nil
}
