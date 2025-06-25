package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
)

// Create create cosmetics and returns it.
func (s *cosmeticsService) Create(ctx context.Context, req *dto.CreateCosmeticsRequest) (*dto.CreateCosmeticsResponse, error) {
	cosmetics := &ent.Cosmetics{
		Category:          req.Category,
		Title:             req.Title,
		Description:       req.Description,
		ApplicationMethod: req.ApplicationMethod,
		Volume:            req.Volume,
	}
	if req.Links != nil {
		cosmetics.OzonLink = req.Links.Ozon
		cosmetics.WildberriesLink = req.Links.Wildberries
	}
	cosmetics, err := s.cosmeticsRepo.Create(ctx, *cosmetics)
	switch {
	case errors.Is(err, errorz.InvalidCosmeticsFormat):
		return nil, errorz.InvalidCosmeticsFormat
	case err != nil:
		return nil, err
	}
	return &dto.CreateCosmeticsResponse{
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
