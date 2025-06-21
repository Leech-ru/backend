package cosmetics

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAllByFilter returns a list of cosmetics based on filter parameters.
//
// @Summary      Get cosmetics by filters
// @Description  Retrieves a list of cosmetics filtered by category, volume, title, etc.
// @Tags         cosmetics
// @Accept       json
// @Produce      json
// @Param        limit        query     int     false  "Max number of items"          minimum(1) maximum(100)
// @Param        offset       query     int     false  "Offset for pagination"        minimum(0)
// @Param        category     query     int     false  "Category enum (0–8)"          Enums(0,1,2,3,4,5,6,7,8)
// @Param        titlePrefix  query     string  false  "Filter by title prefix"
// @Param        volume       query     int     false  "Exact volume in ml"           minimum(1) maximum(10000)
// @Success      200  {array}   dto.Cosmetics
// @Failure      400  {object}  dto.HTTPStatus "Invalid query parameters"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/cosmetics [get]
func (h *handler) GetAllByFilter(c echo.Context) error {
	var req dto.GetAllByFilterCosmeticsRequest

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
	resp, err := h.cosmeticsService.GetAllByFilter(c.Request().Context(), &req)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}
