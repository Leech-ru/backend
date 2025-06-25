package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// GetByID returns the cosmetics by ID.
func (s *cosmeticsService) GetByID(ctx context.Context, req *dto.GetCosmeticsRequest) (*dto.GetCosmeticsResponse, error) {
	cosmetics, err := s.cosmeticsRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		return nil, err
	}

	return &dto.GetCosmeticsResponse{
		ID:                cosmetics.ID,
		Category:          cosmetics.Category,
		Title:             cosmetics.Title,
		Description:       cosmetics.Description,
		ApplicationMethod: cosmetics.ApplicationMethod,
		Volume:            cosmetics.Volume,
		Links: &dto.Links{
			Ozon:        cosmetics.OzonLink,
			Wildberries: cosmetics.WildberriesLink,
		},
	}, nil
}
