package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/partner"
	"Leech-ru/pkg/ent/partnerlink"
	"context"
)

// Update updates an existing partner and returns the partner with loaded links.
func (s *partnersRepo) Update(ctx context.Context, entity ent.Partner) (*ent.Partner, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := tx.PartnerLink.Delete().
		Where(partnerlink.HasPartnerWith(partner.IDEQ(entity.ID))).
		Exec(ctx); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	for _, link := range entity.Edges.Links {
		if _, err := tx.PartnerLink.Create().
			SetLabel(link.Label).
			SetHref(link.Href).
			SetPartnerID(entity.ID).
			Save(ctx); err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}

	_, err = tx.Partner.UpdateOneID(entity.ID).
		SetName(entity.Name).
		SetNillableDescription(entity.Description).
		Save(ctx)

	if err != nil {
		_ = tx.Rollback()
		if ent.IsNotFound(err) {
			return nil, errorz.PartnerNotFound
		}
		return nil, err
	}

	updatedWithLinks, err := tx.Partner.Query().
		Where(partner.IDEQ(entity.ID)).
		WithLinks().
		Only(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return updatedWithLinks, nil
}
