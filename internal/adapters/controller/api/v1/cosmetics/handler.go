package cosmetics

import (
	"Leech-ru/internal/adapters/controller/api/middleware/auth"
	"Leech-ru/internal/adapters/controller/api/middleware/role"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/types"
	"context"
	"github.com/go-playground/form"
	"github.com/labstack/echo/v4"
)

type cosmeticsService interface {
	Create(ctx context.Context, req *dto.CreateCosmeticsRequest) (*dto.CreateCosmeticsResponse, error)
	GetByID(ctx context.Context, req *dto.GetCosmeticsRequest) (*dto.GetCosmeticsResponse, error)
	GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterCosmeticsRequest) (*dto.GetAllByFilterCosmeticsResponse, error)
	Update(ctx context.Context, req *dto.UpdateCosmeticsRequest) (*dto.UpdateCosmeticsResponse, error)
	Delete(ctx context.Context, req *dto.DeleteCosmeticsRequest) error
}

type handler struct {
	cosmeticsService cosmeticsService
	authMiddleware   *auth.Middleware
	roleMiddleware   *role.Middleware
	validator        *validator.Validator
	formDecoder      *form.Decoder
}

func NewHandler(
	cosmeticsService cosmeticsService,
	authMiddleware *auth.Middleware,
	roleMiddleware *role.Middleware,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *handler {
	return &handler{
		cosmeticsService: cosmeticsService,
		authMiddleware:   authMiddleware,
		roleMiddleware:   roleMiddleware,
		validator:        validator,
		formDecoder:      formDecoder,
	}
}
func (h *handler) Setup(router *echo.Group) {
	group := router.Group("/cosmetics", h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))

	group.POST("", h.Create)
	group.GET("/search", h.GetAllByFilter)
	group.GET("/:id", h.GetById)
	group.PATCH("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}
