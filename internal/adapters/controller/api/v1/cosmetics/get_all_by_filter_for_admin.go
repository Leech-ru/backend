package cosmetics

import (
	"Leech-ru/internal/domain/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllByFilterForAdmin returns a list of cosmetics based on filter parameters, but for admin.
//
// @Summary      Get cosmetics by filters
// @Description  Retrieves a list of cosmetics filtered by category, volume, title, etc and hidden parameters.
// @Tags         cosmetics
// @Security     CookieAuth
// @Accept       json
// @Produce      json
// @Param        limit        query     int     false  "Max number of items"          minimum(1) maximum(100)
// @Param        offset       query     int     false  "Offset for pagination"        minimum(0)
// @Param        category_id  query     string  false  "Category UUID"  Format(uuid)  example(123e4567-e89b-12d3-a456-426614174000)
// @Param        titlePrefix  query     string  false  "Filter by title prefix"
// @Param        volume       query     int     false  "Exact volume in ml"           minimum(1) maximum(10000)
// @Success      200  {object}  dto.GetAllByFilterCosmeticsResponse
// @Failure      400  {object}  dto.HTTPStatus "Invalid query parameters"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/cosmetics/admin [get]
func (h *handler) GetAllByFilterForAdmin(c echo.Context) error {
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
