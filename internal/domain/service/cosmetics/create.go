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
	c, err := s.cosmeticsRepo.Create(ctx, ent.Cosmetics{
		Category:          req.Category,
		Title:             req.Title,
		Description:       req.Description,
		ApplicationMethod: req.ApplicationMethod,
		Volume:            req.Volume,
	})
	switch {
	case errors.Is(err, errorz.InvalidCosmeticsFormat):
		return nil, errorz.InvalidCosmeticsFormat
	case err != nil:
		return nil, err
	}
	return &dto.CreateCosmeticsResponse{
		ID:                c.ID,
		Category:          c.Category,
		Title:             c.Title,
		Description:       c.Description,
		ApplicationMethod: c.ApplicationMethod,
		Volume:            c.Volume,
	}, nil
}
