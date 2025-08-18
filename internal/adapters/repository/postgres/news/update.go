package news

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Update updates an existing category
func (s *newsRepo) Update(ctx context.Context, entity ent.News) (*ent.News, error) {
	updated, err := s.client.News.
		UpdateOneID(entity.ID).
		SetImageID(entity.ImageID).
		SetTitle(entity.Title).
		SetContent(entity.Content).
		SetIsHidden(entity.IsHidden).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.NewsNotFound
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidCategoryFormat
	case err != nil:
		return nil, err
	}

	return updated, nil
}
