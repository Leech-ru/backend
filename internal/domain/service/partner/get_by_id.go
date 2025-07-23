package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// GetByID returns partner by ID.
func (s *partnerService) GetByID(ctx context.Context, req *dto.GetByIdPartnerRequest) (*dto.GetByIdPartnerResponse, error) {
	partner, err := s.partnerRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		return nil, err
	}

	return &dto.GetByIdPartnerResponse{
		ID:          partner.ID,
		Name:        partner.Name,
		Description: partner.Description,
	}, nil
}
