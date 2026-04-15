package mainpage

import (
	"Leech-ru/internal/adapters/repository/postgres/main_page"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"

	"github.com/google/uuid"
)

type mainPageRepo interface {
	GetById(ctx context.Context, id uuid.UUID) (*ent.MainPage, error)
	Create(context.Context, ent.MainPage) (*ent.MainPage, error)
	Update(ctx context.Context, entity ent.MainPage) (*ent.MainPage, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context, limit, offset int) ([]*ent.MainPage, int, error)
}

type imageService interface {
	Create(ctx context.Context, req *dto.CreateImageRequest) (*dto.CreateImageResponse, error)
	GetById(ctx context.Context, req *dto.GetByIdImageRequest) (*dto.GetByIdImageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteImageRequest) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

type mainPageService struct {
	mainPageRepo mainPageRepo
	imageService imageService
}

func NewMainPageService(entClient *ent.Client, imageService imageService) *mainPageService {
	return &mainPageService{
		mainPageRepo: mainpage.NewMainPageRepo(entClient),
		imageService: imageService,
	}
}
