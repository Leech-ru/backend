package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Update updates an existing cosmetics
func (s *cosmeticsRepo) Update(ctx context.Context, entity ent.Cosmetics) (*ent.Cosmetics, error) {
	updated, err := s.client.Cosmetics.
		UpdateOneID(entity.ID).
		SetCategoryID(entity.Edges.Category.ID).
		SetTitle(entity.Title).
		SetNillableDescription(entity.Description).
		SetNillableApplicationMethod(entity.ApplicationMethod).
		SetNillableVolume(entity.Volume).
		SetNillableOzonLink(entity.OzonLink).
		SetNillableWildberriesLink(entity.WildberriesLink).
		SetIsHidden(entity.IsHidden).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		return nil, err
	}

	return updated, nil
}
