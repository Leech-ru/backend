package news

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAllByFilterForAdmins returns a list of news with optional filters for admins.
//
// @Summary      Get news by filters (admin)
// @Description  Retrieves a list of news with pagination, preview length option, and hidden flag (for admins).
// @Tags         news
// @Security     CookieAuth
// @Accept       json
// @Produce      json
// @Param        limit           query     int   false  "Max number of items"                  minimum(1) maximum(100)
// @Param        offset          query     int   false  "Offset for pagination"                minimum(0)
// @Param        preview_length  query     int   false  "Number of characters for content preview" minimum(1) maximum(500)
// @Param        is_hidden       query     bool  false  "Filter by hidden status (true/false)"
// @Success      200  {object}  dto.GetAllByFilterForAdminsNewsResponse
// @Failure      400  {object}  dto.HTTPStatus "Invalid query parameters"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/admin/news [get]
func (h *handler) GetAllByFilterForAdmins(c echo.Context) error {
	var req dto.GetAllByFilterForAdminsNewsRequest

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

	resp, err := h.newsService.GetAllByFilterForAdmins(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
