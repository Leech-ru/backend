package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
)

// Create creates a new main page content in the database
func (s *mainPageRepo) Create(ctx context.Context, entity ent.MainPage) (*ent.MainPage, error) {
	created, err := s.client.MainPage.
		Create().
		SetImageID(entity.ImageID).
		SetTitle(entity.Title).
		SetContent(entity.Content).
		SetHref(entity.Href).
		SetIsHidden(entity.IsHidden).
		SetFluid(entity.Fluid).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.InvalidMainPageFormat
	case err != nil:
		return nil, fmt.Errorf("failed to query db: %w", err)
	}
	return created, nil
}
