package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetById returns a cosmetic product by its ID.
//
// @Summary      Get cosmetic by ID
// @Description  Retrieves a cosmetic product using its UUID.
// @Tags         cosmetics
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Cosmetic ID (UUID)"  Format(uuid)
// @Success      200  {object}  dto.GetCosmeticsResponse
// @Failure      400  {object}  dto.HTTPStatus "Validation error"
// @Failure      404  {object}  dto.HTTPStatus "Cosmetic not found"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/cosmetics/{id} [get]
func (h *handler) GetById(c echo.Context) error {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.CosmeticsNotFound.Error(),
		})
	}
	var req dto.GetCosmeticsRequest
	req.ID = userID

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.cosmeticsService.GetByID(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusOK, resp)
}
