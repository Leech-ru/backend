package news

import (
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type newsRepo interface {
	GetById(ctx context.Context, id uuid.UUID) (*ent.News, error)
	Create(context.Context, ent.News) (*ent.News, error)
	Update(ctx context.Context, entity ent.News) (*ent.News, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Get
}
