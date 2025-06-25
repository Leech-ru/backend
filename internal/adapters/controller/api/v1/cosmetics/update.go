package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Update updates a cosmetic product by ID.
//
// @Summary      Update cosmetic
// @Description  Updates cosmetic product fields by given ID.
// @Tags         cosmetics
// @Accept       json
// @Produce      json
// @Param        id      path      string                    true  "Cosmetic ID (UUID)"  Format(uuid)
// @Param        request body      dto.UpdateCosmeticsRequest true  "Updated cosmetic fields"
// @Success      200     {object}  dto.UpdateCosmeticsResponse
// @Failure      400     {object}  dto.HTTPStatus "Validation or binding error"
// @Failure      404     {object}  dto.HTTPStatus "Cosmetic not found"
// @Failure      500     {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/cosmetics/{id} [patch]
func (h *handler) Update(c echo.Context) error {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.CosmeticsNotFound.Error(),
		})
	}
	var req dto.UpdateCosmeticsRequest
	req.ID = userID

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

	resp, err := h.cosmeticsService.Update(c.Request().Context(), &req)
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
