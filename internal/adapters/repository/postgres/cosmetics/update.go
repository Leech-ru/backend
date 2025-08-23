package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/cosmetics"
	"context"
	"fmt"
)

// Update updates an existing cosmetics and returns it with loaded relations
func (s *cosmeticsRepo) Update(ctx context.Context, entity ent.Cosmetics) (*ent.Cosmetics, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin tx: %w", err)
	}

	_, err = tx.Cosmetics.Get(ctx, entity.ID)
	if err != nil {
		_ = tx.Rollback()
		if ent.IsNotFound(err) {
			return nil, errorz.CosmeticsNotFound
		}
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	_, err = tx.Category.Get(ctx, entity.Edges.Category.ID)
	if err != nil {
		_ = tx.Rollback()
		if ent.IsNotFound(err) {
			return nil, errorz.CategoryNotFound
		}
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	_, err = tx.Cosmetics.
		UpdateOneID(entity.ID).
		SetImageID(entity.ImageID).
		SetTitle(entity.Title).
		SetCategoryID(entity.Edges.Category.ID).
		SetNillableDescription(entity.Description).
		SetNillableApplicationMethod(entity.ApplicationMethod).
		SetNillableVolume(entity.Volume).
		SetNillableOzonLink(entity.OzonLink).
		SetNillableWildberriesLink(entity.WildberriesLink).
		SetIsHidden(entity.IsHidden).
		Save(ctx)

	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	result, err := tx.Cosmetics.
		Query().
		Where(cosmetics.ID(entity.ID)).
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
