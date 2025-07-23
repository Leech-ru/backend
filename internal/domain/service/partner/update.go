package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Update change cosmetics data.
func (s *partnerService) Update(ctx context.Context, req *dto.UpdatePartnerRequest) (*dto.UpdatePartnerResponse, error) {
	partnerToUpdate, err := s.partnerRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.PartnerNotFound):
		return nil, errorz.PartnerNotFound
	case err != nil:
		return nil, err
	}
	if req.Name != nil {
		partnerToUpdate.Name = *req.Name
	}
	if req.Description != nil {
		partnerToUpdate.Description = req.Description
	}
	updatedPartner, err := s.partnerRepo.Update(ctx, *partnerToUpdate)
	switch {
	case errors.Is(err, errorz.PartnerNotFound):
		return nil, errorz.PartnerNotFound
	case errors.Is(err, errorz.InvalidPartnerFormat):
		return nil, errorz.InvalidPartnerFormat
	case err != nil:
		return nil, err
	}

	return &dto.UpdatePartnerResponse{
		ID:          updatedPartner.ID,
		Name:        updatedPartner.Name,
		Description: updatedPartner.Description,
	}, nil
}
