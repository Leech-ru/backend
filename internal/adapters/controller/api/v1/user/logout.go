package user

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/internal/domain/utils/cookie"
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
	if err != nil {
		return err
	}
	cookie.ClearAccessTokenCookie(c, h.serverConfig.DevMode())
	cookie.ClearRefreshTokenCookie(c, h.serverConfig.DevMode())

	return c.NoContent(http.StatusNoContent)
}
