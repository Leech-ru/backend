package image

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/image"
	"context"
	"github.com/google/uuid"
)

// GetById retrieves a category by ID
func (s *imageRepo) GetById(ctx context.Context, id uuid.UUID) (*ent.Image, error) {
	c, err := s.client.Image.
		Query().
		Where(image.ID(id)).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.ImageNotFound
	case err != nil:
		return nil, err
	}

	return c, nil
}
