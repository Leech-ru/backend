package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Create handles the creation of a new partner product.
//
// @Summary      Create a new cosmetic product
// @Description  Creates a new cosmetic product with provided details.
// @Tags         partner
// @Security     CookieAuth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreatePartnerRequest  true  "Partner product data"
// @Success      201      {object}  dto.CreatePartnerResponse
// @Failure      400      {object}  dto.HTTPStatus "Invalid request body or validation error"
// @Failure      409      {object}  dto.HTTPStatus "Conflict: invalid partner format"
// @Failure      500      {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/partner [post]
func (h *handler) Create(c echo.Context) error {
	var req dto.CreatePartnerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	resp, err := h.partnerService.Create(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.InvalidPartnerFormat):
		return c.JSON(http.StatusConflict, dto.HTTPStatus{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusCreated, resp)

}
