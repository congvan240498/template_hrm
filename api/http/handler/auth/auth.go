package auth

import (
	"github.com/labstack/echo/v4"
	"hrm/internal/domain"
	"hrm/internal/service"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	req := new(domain.LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.AuthService.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(400, domain.Response{
			Status:  "OK",
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(200, domain.Response{
		Status:  "OK",
		Message: "Login successfully",
		Data:    res,
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	err := h.AuthService.Logout(c.Request().Context(), token)
	if err != nil {
		return c.JSON(400, domain.Response{
			Status:  "OK",
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(200, domain.Response{
		Status:  "OK",
		Message: "Logout successfully",
	})
}
