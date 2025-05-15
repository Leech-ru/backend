package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) ChangePassword(c echo.Context) error {
	var req dto.ChangePasswordRequest
	userID, _ := c.Get("user_id").(uuid.UUID)
	req.ID = userID

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

	resp, err := h.userService.ChangePassword(c.Request().Context(), &req)
	switch {
	case errors.As(err, &errorz.UserNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.UserNotFound.Error(),
		})
	case errors.As(err, &errorz.PasswordMismatch):
		return c.JSON(http.StatusForbidden, dto.HTTPStatus{
			Code:    http.StatusForbidden,
			Message: errorz.PasswordMismatch.Error(),
		})
	case errors.As(err, &errorz.PasswordsCoincidence):
		return c.JSON(http.StatusConflict, dto.HTTPStatus{
			Code:    http.StatusConflict,
			Message: errorz.PasswordsCoincidence.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusCreated, resp)

}
