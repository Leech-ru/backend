package image

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new image in the database
func (s *imageRepo) Create(ctx context.Context, entity ent.Image) (*ent.Image, error) {
	created, err := s.client.Image.Create().
		SetName(entity.Name).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidImageFormat
	case err != nil:
		return nil, err
	}
	return created, nil
}
