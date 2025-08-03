package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new partner in the database with optional links in a transaction.
func (s *partnersRepo) Create(ctx context.Context, entity ent.Partner) (*ent.Partner, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	created, err := tx.Partner.
		Create().
		SetName(entity.Name).
		SetNillableDescription(entity.Description).
		Save(ctx)

	if err != nil {
		_ = tx.Rollback()
		if ent.IsConstraintError(err) {
			return nil, errorz.InvalidPartnerFormat
		}
		return nil, err
	}

	if len(entity.Edges.Links) > 0 {
		builders := make([]*ent.PartnerLinkCreate, 0, len(entity.Edges.Links))
		for _, link := range entity.Edges.Links {
			builders = append(builders, tx.PartnerLink.
				Create().
				SetLabel(link.Label).
				SetHref(link.Href).
				SetPartner(created))
		}

		if err := tx.PartnerLink.CreateBulk(builders...).Exec(ctx); err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return created, nil
}
