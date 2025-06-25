package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Update change cosmetics data.
func (s *cosmeticsService) Update(ctx context.Context, req *dto.UpdateCosmeticsRequest) (*dto.UpdateCosmeticsResponse, error) {
	cosmeticToUpdate, err := s.cosmeticsRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		return nil, err
	}
	if req.Category != nil {
		cosmeticToUpdate.Category = *req.Category
	}
	if req.Title != nil {
		cosmeticToUpdate.Title = *req.Title
	}
	if req.Description != nil {
		cosmeticToUpdate.Description = req.Description
	}
	if req.ApplicationMethod != nil {
		cosmeticToUpdate.ApplicationMethod = req.ApplicationMethod
	}
	if req.Volume != nil {
		cosmeticToUpdate.Volume = req.Volume
	}
	if req.Links != nil {
		if req.Links.Ozon != nil {
			cosmeticToUpdate.OzonLink = req.Links.Ozon
		}
		if req.Links.Wildberries != nil {
			cosmeticToUpdate.WildberriesLink = req.Links.Wildberries
		}
	}

	updatedCosmetic, err := s.cosmeticsRepo.Update(ctx, *cosmeticToUpdate)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return nil, errorz.CosmeticsNotFound
	case errors.Is(err, errorz.InvalidCosmeticsFormat):
		return nil, errorz.InvalidCosmeticsFormat
	case err != nil:
		return nil, err
	}

	return &dto.UpdateCosmeticsResponse{
		ID:                updatedCosmetic.ID,
		Category:          updatedCosmetic.Category,
		Title:             updatedCosmetic.Title,
		Description:       updatedCosmetic.Description,
		ApplicationMethod: updatedCosmetic.ApplicationMethod,
		Volume:            updatedCosmetic.Volume,
		Links: &dto.Links{
			Ozon:        updatedCosmetic.OzonLink,
			Wildberries: updatedCosmetic.WildberriesLink,
		},
	}, nil
}
