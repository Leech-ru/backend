package partner

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll realizes a search for partners with optional pagination.
func (s *partnerService) GetAll(ctx context.Context, req *dto.GetAllPartnerRequest) (*dto.GetAllPartnerResponse, error) {
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}
	allPartners, err := s.partnerRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	var resp dto.GetAllPartnerResponse
	for _, partner := range allPartners {
		resp = append(resp, &dto.Partner{
			ID:          partner.ID,
			Name:        partner.Name,
			Description: partner.Description,
		})
	}
	return &resp, nil
}
