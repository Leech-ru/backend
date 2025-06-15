package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// GetByID returns the cosmetics by ID.
func (s *cosmeticsService) GetByID(ctx context.Context, req *dto.GetCosmeticsRequest) (*dto.GetCosmeticsResponse, error) {
	c, err := s.cosmeticsRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		return nil, err
	}

	return &dto.GetCosmeticsResponse{
		ID:                c.ID,
		Category:          c.Category,
		Title:             c.Title,
		Description:       c.Description,
		ApplicationMethod: c.ApplicationMethod,
		Volume:            c.Volume,
	}, nil
}
