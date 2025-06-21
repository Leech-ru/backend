package cosmetics

import (
	"Leech-ru/internal/adapters/controller/api/middleware/auth"
	"Leech-ru/internal/adapters/controller/api/middleware/role"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/internal/domain/dto"
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
	router.GET("/cosmetics/search", h.GetAllByFilter)
	router.GET("/cosmetics/:id", h.GetById)
	router.POST("/cosmetics", h.Create)
	router.PATCH("/cosmetics/:id", h.Update)
	router.DELETE("/cosmetics/:id", h.Delete)

	//TODO включить когда будут юзеры
	//router.POST("/cosmetics", h.Create, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	//router.PATCH("/cosmetics/:id", h.Update, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	//router.DELETE("/cosmetics/:id", h.Delete, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
}
