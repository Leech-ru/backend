package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/cosmetics"
	"context"
)

// Update updates an existing cosmetics and returns it with loaded relations
func (s *cosmeticsRepo) Update(ctx context.Context, entity ent.Cosmetics) (*ent.Cosmetics, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.Cosmetics.
		UpdateOneID(entity.ID).
		SetTitle(entity.Title).
		SetCategoryID(entity.Edges.Category.ID).
		SetNillableDescription(entity.Description).
		SetNillableApplicationMethod(entity.ApplicationMethod).
		SetNillableVolume(entity.Volume).
		SetNillableOzonLink(entity.OzonLink).
		SetNillableWildberriesLink(entity.WildberriesLink).
		SetIsHidden(entity.IsHidden).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		_ = tx.Rollback()
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		_ = tx.Rollback()
		return nil, err
	}

	result, err := tx.Cosmetics.
		Query().
		Where(cosmetics.ID(entity.ID)).
		WithCategory().
		Only(ctx)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}
