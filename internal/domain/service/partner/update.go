package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Update changes partner data.
func (s *partnerService) Update(ctx context.Context, req *dto.UpdatePartnerRequest) (*dto.UpdatePartnerResponse, error) {
	partnerToUpdate, err := s.partnerRepo.GetById(ctx, req.ID)
	if err != nil {
		if errors.Is(err, errorz.PartnerNotFound) {
			return nil, errorz.PartnerNotFound
		}
		return nil, err
	}

	if req.Name != nil {
		partnerToUpdate.Name = *req.Name
	}
	if req.Description != nil {
		partnerToUpdate.Description = req.Description
	}

	if req.Links != nil {
		partnerToUpdate.Edges.Links = convertDtoLinksToEnt(req.Links)
	}

	updatedPartner, err := s.partnerRepo.Update(ctx, *partnerToUpdate)
	if err != nil {
		switch {
		case errors.Is(err, errorz.PartnerNotFound):
			return nil, errorz.PartnerNotFound
		case errors.Is(err, errorz.InvalidPartnerFormat):
			return nil, errorz.InvalidPartnerFormat
		default:
			return nil, err
		}
	}

	return &dto.UpdatePartnerResponse{
		ID:          updatedPartner.ID,
		Name:        updatedPartner.Name,
		Description: updatedPartner.Description,
		Links:       convertLinksToDto(updatedPartner.Edges.Links),
	}, nil
}
