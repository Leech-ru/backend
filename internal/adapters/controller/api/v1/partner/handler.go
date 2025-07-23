package partner

import (
	"Leech-ru/internal/adapters/controller/api/middleware/auth"
	"Leech-ru/internal/adapters/controller/api/middleware/role"
	"Leech-ru/internal/adapters/controller/api/validator"
	"Leech-ru/internal/domain/dto"
	"context"
	"github.com/go-playground/form"
	"github.com/labstack/echo/v4"
)

type partnerService interface {
	Create(ctx context.Context, req *dto.CreatePartnerRequest) (*dto.CreatePartnerResponse, error)
	GetByID(ctx context.Context, req *dto.GetByIdPartnerRequest) (*dto.GetByIdPartnerResponse, error)
	GetAll(ctx context.Context, req *dto.GetAllPartnerRequest) (*dto.GetAllPartnerResponse, error)
	Update(ctx context.Context, req *dto.UpdatePartnerRequest) (*dto.UpdatePartnerResponse, error)
	Delete(ctx context.Context, req *dto.DeletePartnerRequest) error
}

type handler struct {
	partnerService partnerService
	authMiddleware *auth.Middleware
	roleMiddleware *role.Middleware
	validator      *validator.Validator
	formDecoder    *form.Decoder
}

func NewHandler(
	partnerService partnerService,
	authMiddleware *auth.Middleware,
	roleMiddleware *role.Middleware,
	validator *validator.Validator,
	formDecoder *form.Decoder,

) *handler {
	return &handler{
		partnerService: partnerService,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
		validator:      validator,
		formDecoder:    formDecoder,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.GET("/partner", h.GetAll)
	router.GET("/partner/:id", h.GetById)
	router.POST("/partner", h.Create)
	router.PATCH("/partner/:id", h.Update)
	router.DELETE("/partner/:id", h.Delete)

}
