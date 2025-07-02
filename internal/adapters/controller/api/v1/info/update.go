package info

import (
	"Leech-ru/internal/domain/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

// UpdateInfo обновляет информацию о корпорации.
//
// @Summary      Update corporation info
// @Description  Updates the corporation info fields. Only accessible by authenticated moderators.
// @Description  Requires authentication via cookies (access_token, refresh_token)
// @Tags         info
// @Accept       json
// @Produce      json
// @Param        request   body  dto.UpdateInfoRequest  true  "Updated corporation info"
// @Success      200   {object}  dto.UpdateInfoResponse
// @Failure      400   {object}  dto.HTTPStatus "Validation error"
// @Failure      401   {object}  dto.HTTPStatus "Unauthorized"
// @Failure      403   {object}  dto.HTTPStatus "Forbidden"
// @Failure      500   {object}  dto.HTTPStatus "Internal server error"
// @Router       /info/corporation [patch]
func (h *handler) UpdateInfo(c echo.Context) error {
	var req dto.UpdateInfoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := h.validator.ValidateData(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := h.infoService.Update(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
