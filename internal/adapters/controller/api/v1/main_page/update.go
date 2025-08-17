package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Update updates a main page content by ID.
//
// @Summary      Update content
// @Description  Updates main page content fields by given ID.
// @Tags         main
// @Security CookieAuth
// @Accept       json
// @Produce      json
// @Param        id      path      string                    true  "Content ID (UUID)"  Format(uuid)
// @Param        request body      dto.UpdateMainPageRequest true  "Updated content fields"
// @Success      200     {object}  dto.UpdateMainPageResponse
// @Failure      400     {object}  dto.HTTPStatus "Validation or binding error"
// @Failure      404     {object}  dto.HTTPStatus "Content not found"
// @Failure      500     {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/main/{id} [patch]
func (h *handler) Update(c echo.Context) error {
	id := c.Param("id")
	pageID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.MainPageNotFound.Error(),
		})
	}

	var req dto.UpdateMainPageRequest
	req.ID = pageID

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

	resp, err := h.mainPageService.Update(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.MainPageNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case errors.Is(err, errorz.ImageNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case errors.Is(err, errorz.InvalidMainPageFormat):
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

	return c.JSON(http.StatusOK, resp)
}
