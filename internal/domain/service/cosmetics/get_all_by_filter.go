package cosmetics

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAllByFilter realizes a search for cosmetics with filtering parameters.
func (s *cosmeticsService) GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterCosmeticsRequest) (*dto.GetAllByFilterCosmeticsResponse, error) {
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}
	isHidden := false
	allCosmetics, err := s.cosmeticsRepo.GetAllByFilter(ctx, limit, offset, req.CategoryID, req.TitlePrefix, req.Volume, &isHidden)
	if err != nil {
		return nil, err
	}
	var resp dto.GetAllByFilterCosmeticsResponse
	for _, cosmetics := range allCosmetics {
		resp = append(resp, &dto.Cosmetics{
			ID: cosmetics.ID,
			Category: dto.Category{
				ID:   cosmetics.Edges.Category.ID,
				Name: cosmetics.Edges.Category.Name,
			},
			Title:             cosmetics.Title,
			Description:       cosmetics.Description,
			ApplicationMethod: cosmetics.ApplicationMethod,
			Volume:            cosmetics.Volume,
			Links: &dto.Links{
				Ozon:        cosmetics.OzonLink,
				Wildberries: cosmetics.WildberriesLink,
			},
			IsHidden: cosmetics.IsHidden,
			ImageID:  cosmetics.ImageID,
		})
	}
	return &resp, nil
}
