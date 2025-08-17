package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Update updates an existing main page
func (s *mainPageRepo) Update(ctx context.Context, entity ent.MainPage) (*ent.MainPage, error) {
	updated, err := s.client.MainPage.
		UpdateOneID(entity.ID).
		SetImageID(entity.ImageID).
		SetTitle(entity.Title).
		SetContent(entity.Content).
		SetHref(entity.Href).
		SetIsHidden(entity.IsHidden).
		SetFluid(entity.Fluid).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, errorz.MainPageNotFound
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidMainPageFormat
	case err != nil:
		return nil, err
	}

	return updated, nil
}
