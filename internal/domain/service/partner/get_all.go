package partner

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll реализует поиск партнеров с опциональной пагинацией.
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

	resp := make(dto.GetAllPartnerResponse, 0, len(allPartners))
	for _, partner := range allPartners {
		resp = append(resp, &dto.Partner{
			ID:          partner.ID,
			Name:        partner.Name,
			Description: partner.Description,
			Links:       convertLinksToDto(partner.Edges.Links),
		})
	}

	return &resp, nil
}
