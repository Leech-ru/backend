package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new cosmetics in the database
func (s *cosmeticsRepo) Create(ctx context.Context, entity ent.Cosmetics) (*ent.Cosmetics, error) {
	created, err := s.client.Cosmetics.
		Create().
		SetCategory(entity.Category).
		SetTitle(entity.Title).
		SetNillableDescription(entity.Description).
		SetNillableApplicationMethod(entity.ApplicationMethod).
		SetNillableVolume(entity.Volume).
		SetNillableOzonLink(entity.OzonLink).
		SetNillableWildberriesLink(entity.WildberriesLink).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidCosmeticsFormat
	case err != nil:
		return nil, err
	}
	return created, nil
}
