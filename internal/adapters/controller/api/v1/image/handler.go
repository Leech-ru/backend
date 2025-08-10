package image

import (
	"Leech-ru/internal/adapters/controller/api/middleware/auth"
	"Leech-ru/internal/adapters/controller/api/middleware/role"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/types"
	"context"
	"github.com/labstack/echo/v4"
)

type imageService interface {
	Create(ctx context.Context, req *dto.CreateImageRequest) (*dto.CreateImageResponse, error)
	GetById(ctx context.Context, req *dto.GetByIdImageRequest) (*dto.GetByIdImageResponse, error)
	Delete(ctx context.Context, req *dto.DeleteImageRequest) error
}

type handler struct {
	imageService   imageService
	authMiddleware *auth.Middleware
	roleMiddleware *role.Middleware
	validator      *validator.Validator
}

func NewHandler(
	imageService imageService,
	authMiddleware *auth.Middleware,
	roleMiddleware *role.Middleware,
	validator *validator.Validator,

) *handler {
	return &handler{
		imageService:   imageService,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
		validator:      validator,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.GET("/image/:image_id", h.Download)
	router.POST("/image", h.Upload, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.DELETE("/image/:image_id", h.Delete, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
}
