package cosmetics

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAllByFilter realizes a search for cosmetics with filtering parameters.
func (s *cosmeticsService) GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterCosmeticsRequest) (*dto.GetAllByFilterCosmeticsResponse, error) {
	limit := 10
	if req.Limit != nil && *req.Limit > 0 {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}
	isHidden := false
	allCosmetics, totalItems, err := s.cosmeticsRepo.GetAllByFilter(ctx, limit, offset, req.CategoryID, req.TitlePrefix, req.Volume, &isHidden)
	if err != nil {
		return nil, err
	}
	respItems := make([]*dto.Cosmetics, 0, len(allCosmetics))
	for _, cosmetics := range allCosmetics {
		respItems = append(respItems, &dto.Cosmetics{
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

	totalPages := 0
	if totalItems > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}
	currentPage := (offset / limit) + 1

	resp := dto.GetAllByFilterCosmeticsResponse{
		Items: respItems,
		Pagination: dto.PaginationInfo{
			TotalItems:  totalItems,
			TotalPages:  totalPages,
			CurrentPage: currentPage,
			HasNext:     currentPage < totalPages,
			HasPrevious: currentPage > 1,
		},
	}

	return &resp, nil
}
