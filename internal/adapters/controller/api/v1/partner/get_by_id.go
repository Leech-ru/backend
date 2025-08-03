package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetById returns a partner product by its ID.
//
// @Summary      Get partner by ID
// @Description  Retrieves a partner product using its UUID.
// @Tags         partner
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Partner ID (UUID)"  Format(uuid)
// @Success      200  {object}  dto.GetByIdPartnerResponse
// @Failure      400  {object}  dto.HTTPStatus "Validation error"
// @Failure      404  {object}  dto.HTTPStatus "Partner not found"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/info/partner/{id} [get]
func (h *handler) GetById(c echo.Context) error {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.CosmeticsNotFound.Error(),
		})
	}
	var req dto.GetByIdPartnerRequest
	req.ID = userID

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.partnerService.GetByID(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.PartnerNotFound):
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
