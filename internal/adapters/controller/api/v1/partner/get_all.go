package partner

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAll returns a list of cosmetics based on filter parameters.
//
// @Summary      Get partners with pagination
// @Description  Retrieves a list of partners with pagination.
// @Tags         partner
// @Accept       json
// @Produce      json
// @Param        limit        query     int     false  "Max number of items"          minimum(1) maximum(100)
// @Param        offset       query     int     false  "Offset for pagination"        minimum(0)
// @Success      200  {array}   dto.Partner
// @Failure      400  {object}  dto.HTTPStatus "Invalid query parameters"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/partner [get]
func (h *handler) GetAll(c echo.Context) error {
	var req dto.GetAllPartnerRequest

	if err := h.formDecoder.Decode(&req, c.QueryParams()); err != nil {
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
	resp, err := h.partnerService.GetAll(c.Request().Context(), &req)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}
