package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/utils/cookie"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
	cookie.ClearAccessTokenCookie(c, h.serverConfig.DevMode())
	cookie.ClearRefreshTokenCookie(c, h.serverConfig.DevMode())

	return c.NoContent(http.StatusNoContent)
}
