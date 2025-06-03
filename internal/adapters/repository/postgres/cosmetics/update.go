package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Update updates an existing cosmetics
func (s *cosmeticsRepo) Update(ctx context.Context, userEntity ent.Cosmetics) (*ent.Cosmetics, error) {
	updated, err := s.client.Cosmetics.
		UpdateOneID(userEntity.ID).
		SetCategory(userEntity.Category).
		SetTitle(userEntity.Title).
		SetNillableDescription(userEntity.Description).
		SetNillableApplicationMethod(userEntity.ApplicationMethod).
		SetNillableVolume(userEntity.Volume).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		return nil, err
	}

	return updated, nil
}
