package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new category in the database
func (s *newsRepo) Create(ctx context.Context, entity ent.News) (*ent.News, error) {
	created, err := s.client.News.
		Create().
		SetImageID(entity.ImageID).
		SetTitle(entity.Title).
		SetContent(entity.Content).
		SetIsHidden(entity.IsHidden).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidNewsFormat
	case err != nil:
		return nil, err
	}
	return created, nil
}
