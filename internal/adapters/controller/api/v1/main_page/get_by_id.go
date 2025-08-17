package mainpage

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetById returns a main page content by its ID.
//
// @Summary      Get content by ID
// @Description  Retrieves a main page content using its UUID.
// @Tags         main
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Content ID (UUID)"  Format(uuid)
// @Success      200  {object}  dto.GetByIdMainPageResponse
// @Failure      400  {object}  dto.HTTPStatus "Validation error"
// @Failure      404  {object}  dto.HTTPStatus "Content not found"
// @Failure      500  {object}  dto.HTTPStatus "Internal server error"
// @Router       /api/v1/main/{id} [get]
func (h *handler) GetById(c echo.Context) error {
	id := c.Param("id")
	pageID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.MainPageNotFound.Error(),
		})
	}

	var req dto.GetByIdMainPageRequest
	req.ID = pageID

	if err = h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.mainPageService.GetByID(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.MainPageNotFound):
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
