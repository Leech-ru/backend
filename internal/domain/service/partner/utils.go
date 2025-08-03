package partner

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
)

func convertLinksToDto(entLinks []*ent.PartnerLink) []dto.PartnersLink {
	if entLinks == nil {
		return nil
	}
	links := make([]dto.PartnersLink, 0, len(entLinks))
	for _, l := range entLinks {
		links = append(links, dto.PartnersLink{
			Label: l.Label,
			Href:  l.Href,
		})
	}
	return links
}

func convertDtoLinksToEnt(dtoLinks []*dto.PartnersLink) []*ent.PartnerLink {
	var entLinks []*ent.PartnerLink
	for _, l := range dtoLinks {
		entLinks = append(entLinks, &ent.PartnerLink{
			Label: l.Label,
			Href:  l.Href,
		})
	}
	return entLinks
}
