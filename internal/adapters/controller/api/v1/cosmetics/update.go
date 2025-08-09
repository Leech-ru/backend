package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Update updates a cosmetic product by ID.
//
// @Summary      Update cosmetic product by ID
// @Description  Updates cosmetic product fields by given ID.
//               Accepts multipart/form-data with JSON in "payload" field and optional image file in "file" field.
// @Tags         cosmetics
// @Security     CookieAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id      path      string                    true  "Cosmetic ID (UUID)"  Format(uuid)
// @Param        payload formData  string                    true  "JSON string with updated cosmetic fields"
// @Param        file    formData  file                      false "Optional image file"
// @Success      200     {object}  dto.UpdateCosmeticsResponse
// @Failure      400     {object}  dto.HTTPStatus           "Validation or binding error"
// @Failure      404     {object}  dto.HTTPStatus           "Cosmetic, category or image not found"
// @Failure      500     {object}  dto.HTTPStatus           "Internal server error"
// @Router       /api/v1/cosmetics/{id} [patch]

func (h *handler) Update(c echo.Context) error {
	id := c.Param("id")
	cosmeticID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.CosmeticsNotFound.Error(),
		})
	}

	if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Получаем JSON из поля payload
	jsonStr := c.FormValue("payload")
	if jsonStr == "" {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: "payload is required",
		})
	}

	var req dto.UpdateCosmeticsRequest
	if err := json.Unmarshal([]byte(jsonStr), &req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	req.ID = cosmeticID

	// Валидация
	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Получаем файл, если есть (файл необязательный)
	fileHeader, err := c.FormFile("file")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Вызываем сервис с передачей req и файла
	resp, err := h.cosmeticsService.Update(c.Request().Context(), &req, fileHeader)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case errors.Is(err, errorz.CategoryNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
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

	return c.JSON(http.StatusOK, resp)
}
