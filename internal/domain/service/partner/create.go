package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
	"fmt"
)

// Create partner and returns it.
func (s *partnerService) Create(ctx context.Context, req *dto.CreatePartnerRequest) (*dto.CreatePartnerResponse, error) {
	partner := &ent.Partner{
		Name:        req.Name,
		Description: req.Description,
	}
	partner, err := s.partnerRepo.Create(ctx, *partner)
	switch {
	case errors.Is(err, errorz.InvalidPartnerFormat):
		return nil, errorz.InvalidPartnerFormat
	case err != nil:
		return nil, err
	}
	fmt.Println(partner)
	fmt.Println(partner.Description)
	return &dto.CreatePartnerResponse{
		ID:          partner.ID,
		Name:        partner.Name,
		Description: partner.Description,
	}, nil
}
