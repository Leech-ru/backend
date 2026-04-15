package partner

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll реализует поиск партнеров с опциональной пагинацией.
func (s *partnerService) GetAll(ctx context.Context, req *dto.GetAllPartnerRequest) (*dto.GetAllPartnerResponse, error) {
	limit := 10
	if req.Limit != nil && *req.Limit > 0 {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}

	allPartners, totalItems, err := s.partnerRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	respItems := make([]*dto.Partner, 0, len(allPartners))
	for _, partner := range allPartners {
		respItems = append(respItems, &dto.Partner{
			ID:          partner.ID,
			Name:        partner.Name,
			Description: partner.Description,
			Links:       convertLinksToDto(partner.Edges.Links),
		})
	}

	totalPages := 0
	if totalItems > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}
	currentPage := (offset / limit) + 1

	resp := dto.GetAllPartnerResponse{
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
