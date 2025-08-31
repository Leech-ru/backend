package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/cosmetics"
	"context"
	"fmt"
)

// Create creates a new cosmetics in the database with loaded relations in a single transaction
func (s *cosmeticsRepo) Create(ctx context.Context, entity ent.Cosmetics) (*ent.Cosmetics, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin tx: %w", err)
	}

	_, err = tx.Category.Get(ctx, entity.Edges.Category.ID)
	if err != nil {
		_ = tx.Rollback()
		if ent.IsNotFound(err) {
			return nil, errorz.CategoryNotFound
		}
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	query :=
		tx.Cosmetics.
			Create().
			SetTitle(entity.Title).
			SetImageID(entity.ImageID).
			SetCategoryID(entity.Edges.Category.ID).
			SetNillableDescription(entity.Description).
			SetNillableApplicationMethod(entity.ApplicationMethod).
			SetNillableVolume(entity.Volume).
			SetNillableOzonLink(entity.OzonLink).
			SetNillableWildberriesLink(entity.WildberriesLink).
			SetIsHidden(entity.IsHidden)
	created, err := query.Save(ctx)

	if err != nil {
		_ = tx.Rollback()
		if ent.IsConstraintError(err) {
			fmt.Println(err)
			return nil, errorz.InvalidCosmeticsFormat
		}
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	result, err := tx.Cosmetics.
		Query().
		Where(cosmetics.ID(created.ID)).
		WithCategory().
		Only(ctx)

	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit tx: %w", err)
	}

	return result, nil
}
