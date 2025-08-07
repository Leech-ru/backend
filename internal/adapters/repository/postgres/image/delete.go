package image

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

// Delete removes an image by ID
func (s *imageRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.client.Image.
		DeleteOneID(id).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return errorz.ImageNotFound
	case err != nil:
		return err
	}

	return nil
}
