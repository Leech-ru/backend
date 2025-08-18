package image

import (
	"Leech-ru/internal/domain/dto"
	"bytes"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

// UploadImage godoc
// @Summary      Upload a new image
// @Description  Uploads an image file and stores it in the system
// @Tags         image
// @Security     CookieAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "Image file"
// @Success      201   {object}  dto.CreateImageResponse
// @Failure      400   {object}  dto.HTTPStatus  "Invalid request or file is missing"
// @Failure      500   {object}  dto.HTTPStatus  "Internal server error"
// @Router       /images [post]
func (h *handler) UploadImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: "file is required",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	defer src.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	req := dto.CreateImageRequest{
		File: &dto.FilePackage{
			Content:     bytes.NewReader(buf.Bytes()),
			ContentType: file.Header.Get("Content-Type"),
			Size:        file.Size,
			Filename:    file.Filename,
		},
	}

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.imageService.Create(c.Request().Context(), &req)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, resp)
}
