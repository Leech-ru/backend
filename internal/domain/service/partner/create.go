package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"errors"
)

// Create partner and returns it.
func (s *partnerService) Create(ctx context.Context, req *dto.CreatePartnerRequest) (*dto.CreatePartnerResponse, error) {
	entity := ent.Partner{
		Name:        req.Name,
		Description: req.Description,
	}

	if len(req.Links) > 0 {
		entity.Edges.Links = make([]*ent.PartnerLink, 0, len(req.Links))
		for _, link := range req.Links {
			entity.Edges.Links = append(entity.Edges.Links, &ent.PartnerLink{
				Label: link.Label,
				Href:  link.Href,
			})
		}
	}

	created, err := s.partnerRepo.Create(ctx, entity)
	switch {
	case errors.Is(err, errorz.InvalidPartnerFormat):
		return nil, errorz.InvalidPartnerFormat
	case err != nil:
		return nil, err
	}

	resp := &dto.CreatePartnerResponse{
		ID:          created.ID,
		Name:        created.Name,
		Description: created.Description,
	}

	if len(entity.Edges.Links) > 0 {
		resp.Links = make([]dto.PartnersLink, 0, len(entity.Edges.Links))
		for _, link := range entity.Edges.Links {
			if link == nil {
				continue
			}
			resp.Links = append(resp.Links, dto.PartnersLink{
				Label: link.Label,
				Href:  link.Href,
			})
		}
	}

	return resp, nil
}
