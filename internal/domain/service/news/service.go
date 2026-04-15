package news

import (
	"Leech-ru/internal/adapters/repository/postgres/news"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
)

type newsRepo interface {
	Create(context.Context, ent.News) (*ent.News, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.News, error)
	GetAllByFilter(ctx context.Context, limit, offset int, isHidden *bool) ([]*ent.News, int, error)
	Update(ctx context.Context, entity ent.News) (*ent.News, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type imageService interface {
	Create(ctx context.Context, req *dto.CreateImageRequest) (*dto.CreateImageResponse, error)
	GetById(ctx context.Context, req *dto.GetByIdImageRequest) (*dto.GetByIdImageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteImageRequest) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

const basePreviewLength = 50

type newsService struct {
	newsRepo     newsRepo
	imageService imageService
}

func NewNewsService(entClient *ent.Client, imageService imageService) *newsService {
	return &newsService{
		newsRepo:     news.NewNewsRepo(entClient),
		imageService: imageService,
	}
}
