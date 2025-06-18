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
	allCosmetics, err := s.cosmeticsRepo.GetAllByFilter(ctx, limit, offset, req.Category, req.TitlePrefix, req.Limit)
	if err != nil {
		return nil, err
	}
	var resp dto.GetAllByFilterCosmeticsResponse
	for _, cosmetics := range allCosmetics {
		resp = append(resp, &dto.Cosmetics{
			ID:                cosmetics.ID,
			Category:          cosmetics.Category,
			Title:             cosmetics.Title,
			Description:       cosmetics.Description,
			ApplicationMethod: cosmetics.ApplicationMethod,
			Volume:            cosmetics.Volume,
		})
	}
	return &resp, nil
}
