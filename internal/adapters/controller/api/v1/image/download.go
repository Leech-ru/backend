package image

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// StreamImage godoc
// @Summary      Download image by ID
// @Description  Returns the image file stream by its UUID identifier
// @Tags         image
// @Accept       json
// @Produce      octet-stream
// @Param        image_id  path      string  true  "Image UUID"
// @Success      200       {file}    binary  "Image file stream"
// @Failure      400       {object}  dto.HTTPStatus  "Invalid request"
// @Failure      404       {object}  dto.HTTPStatus  "Image not found"
// @Failure      500       {object}  dto.HTTPStatus  "Internal server error"
// @Router       /image/{image_id} [get]
func (h *handler) StreamImage(c echo.Context) error {
	id := c.Param("image_id")
	imageID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.ImageNotFound.Error(),
		})
	}
	var req dto.GetByIdImageRequest
	req.ID = imageID

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.imageService.GetById(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.ImageNotFound):
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

	return c.Stream(http.StatusOK, resp.File.ContentType, resp.File.Content)
}
