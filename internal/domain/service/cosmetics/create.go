package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
)

// Create create cosmetics and returns it.
func (s *cosmeticsService) Create(ctx context.Context, req *dto.CreateCosmeticsRequest) (*dto.CreateCosmeticsResponse, error) {
	cosmetics := &ent.Cosmetics{
		ImageID:           req.ImageID,
		Title:             req.Title,
		Description:       req.Description,
		ApplicationMethod: req.ApplicationMethod,
		Volume:            req.Volume,
		IsHidden:          *req.IsHidden,
		Edges: ent.CosmeticsEdges{
			Category: &ent.Category{ID: req.CategoryID},
		},
	}

	if req.Links != nil {
		cosmetics.OzonLink = req.Links.Ozon
		cosmetics.WildberriesLink = req.Links.Wildberries
	}

	if req.ImageID != nil {
		if cond, err := s.imageService.Exists(ctx, *req.ImageID); err != nil || !cond {
			switch {
			case !cond:
				return nil, errorz.ImageNotFound
			case err != nil:
				return nil, err
			}
		}
	}
	cosmetics, err := s.cosmeticsRepo.Create(ctx, *cosmetics)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCosmeticsResponse{
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
	}, nil
}
