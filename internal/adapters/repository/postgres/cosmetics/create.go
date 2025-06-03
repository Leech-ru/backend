package cosmetics

import (
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
		Save(ctx)

	switch {
	case err != nil:
		return nil, err
	}
	return &ent.Cosmetics{
		ID:                created.ID,
		Category:          created.Category,
		Title:             created.Title,
		Description:       created.Description,
		ApplicationMethod: created.ApplicationMethod,
		Volume:            created.Volume,
	}, nil
}
