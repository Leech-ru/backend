package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
)

// Create partner and returns it.
func (s *partnerService) Create(ctx context.Context, req *dto.CreatePartnerRequest) (*dto.CreatePartnerRequest, error) {
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
	return &dto.CreatePartnerRequest{
		Name:        partner.Name,
		Description: partner.Description,
	}, nil
}
