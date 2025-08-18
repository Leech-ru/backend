package news

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

type newsService interface {
	Create(ctx context.Context, req *dto.CreateNewsRequest) (*dto.CreateNewsResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdNewsRequest) (*dto.GetByIdNewsResponse, error)
	GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterNewsRequest) (*dto.GetAllByFilterNewsResponse, error)
	GetAllByFilterForAdmins(ctx context.Context, req *dto.GetAllByFilterForAdminsNewsRequest) (*dto.GetAllByFilterForAdminsNewsResponse, error)
	Update(ctx context.Context, req *dto.UpdateNewsRequest) (*dto.UpdateNewsResponse, error)
	Delete(ctx context.Context, req *dto.DeleteNewsRequest) error
}

type handler struct {
	newsService    newsService
	authMiddleware *auth.Middleware
	roleMiddleware *role.Middleware
	validator      *validator.Validator
	formDecoder    *form.Decoder
}

func NewHandler(
	newsService newsService,
	authMiddleware *auth.Middleware,
	roleMiddleware *role.Middleware,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *handler {
	return &handler{
		newsService:    newsService,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
		validator:      validator,
		formDecoder:    formDecoder,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.GET("/news", h.GetAllByFilter)
	router.GET("/news/admin", h.GetAllByFilterForAdmins, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.GET("/news/:id", h.GetById)
	router.POST("/news", h.Create, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.PATCH("/news/:id", h.Update, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
	router.DELETE("/news/:id", h.Delete, h.authMiddleware.RequireAuth, h.roleMiddleware.RequireRole(types.RoleModerator))
}
