package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Logout User logout
//
// @Summary Logout user
// @Tags user
// @Accept json
// @Produce json
// @Security CookieAuth
// @Success 204
// @Header 204 {string} Set-Cookie "user_auth_access_token=; expires=Thu, 01 Jan 1970 00:00:00 GMT; Path=/; HttpOnly"
// @Header 204 {string} Set-Cookie "user_auth_refresh_token=; expires=Thu, 01 Jan 1970 00:00:00 GMT; Path=/; HttpOnly"
// @Failure 400 {object} dto.HTTPStatus
// @Router /api/v1/user/logout [post]
func (h *handler) Logout(c echo.Context) error {
	var req dto.LogoutRequest
	userID, _ := c.Get("user_id").(uuid.UUID)
	req.ID = userID

	if err := h.validator.ValidateData(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := h.userService.Logout(c.Request().Context(), &req)
	switch {
	case errors.Is(err, errorz.Unauthorized):
		return c.JSON(http.StatusForbidden, dto.HTTPStatus{
			Code:    http.StatusForbidden,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	h.cookieService.ClearAccessTokenCookie(c, h.serverConfig.DevMode())
	h.cookieService.ClearRefreshTokenCookie(c, h.serverConfig.DevMode())

	return c.NoContent(http.StatusNoContent)
}
