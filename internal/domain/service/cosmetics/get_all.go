package cosmetics

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll returns list of cosmetics.
func (s *cosmeticsService) GetAll(ctx context.Context, req *dto.GetAllCosmeticsRequest) (*dto.GetAllCosmeticsResponse, error) {
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}

	cosmetics, err := s.cosmeticsRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var cosmeticsList dto.GetAllCosmeticsResponse
	for _, cosmetic := range cosmetics {
		cosmeticsList = append(cosmeticsList, dto.Cosmetics{
			ID:                cosmetic.ID,
			Category:          cosmetic.Category,
			Title:             cosmetic.Title,
			Description:       cosmetic.Description,
			ApplicationMethod: cosmetic.ApplicationMethod,
			Volume:            cosmetic.Volume,
		})
	}

	return &cosmeticsList, nil
}
