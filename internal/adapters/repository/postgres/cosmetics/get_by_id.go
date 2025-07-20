package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/cosmetics"
	"context"
	"github.com/google/uuid"
)

// GetById retrieves a cosmetics by ID
func (s *cosmeticsRepo) GetById(ctx context.Context, id uuid.UUID) (*ent.Cosmetics, error) {
	c, err := s.client.Cosmetics.
		Query().
		Where(cosmetics.ID(id)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		return nil, err
	}

	return c, nil
}
