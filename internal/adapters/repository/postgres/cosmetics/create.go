package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/cosmetics"
	"context"
	"fmt"
	"github.com/google/uuid"
)

// Create creates a new cosmetics in the database with loaded relations in a single transaction
func (s *cosmeticsRepo) Create(ctx context.Context, entity ent.Cosmetics) (*ent.Cosmetics, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.Category.Get(ctx, entity.Edges.Category.ID)
	if err != nil {
		_ = tx.Rollback()
		if ent.IsNotFound(err) {
			return nil, errorz.CategoryNotFound
		}
		return nil, err
	}

	if entity.Edges.Images != nil && entity.Edges.Images.ID != uuid.Nil {
		_, err = tx.Image.Get(ctx, entity.Edges.Images.ID)
		if err != nil {
			_ = tx.Rollback()
			if ent.IsNotFound(err) {
				return nil, errorz.ImageNotFound
			}
			return nil, err
		}
	}

	query :=
		tx.Cosmetics.
			Create().
			SetTitle(entity.Title).
			SetCategoryID(entity.Edges.Category.ID).
			SetNillableDescription(entity.Description).
			SetNillableApplicationMethod(entity.ApplicationMethod).
			SetNillableVolume(entity.Volume).
			SetNillableOzonLink(entity.OzonLink).
			SetNillableWildberriesLink(entity.WildberriesLink).
			SetIsHidden(entity.IsHidden)
	if entity.Edges.Images != nil && entity.Edges.Images.ID != uuid.Nil {
		query = query.SetImagesID(entity.Edges.Images.ID)
	}
	created, err := query.Save(ctx)

	if err != nil {
		_ = tx.Rollback()
		if ent.IsConstraintError(err) {
			fmt.Println(err)
			return nil, errorz.InvalidCosmeticsFormat
		}
		return nil, err
	}

	result, err := tx.Cosmetics.
		Query().
		Where(cosmetics.ID(created.ID)).
		WithCategory().
		Only(ctx)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}
