package user

import (
	"github.com/labstack/echo/v4"
	"hrm/internal/domain"
	"hrm/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	req := new(domain.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, domain.Response{
			Status:  "OK",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err := h.userService.CreateUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(400, domain.Response{
			Status:  "OK",
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(200, domain.Response{
		Status:  "OK",
		Message: "Create user successfully",
	})
}

func (h *UserHandler) GetUser(c echo.Context) error {
	req := new(domain.GetUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, domain.Response{
			Status:  "OK",
			Message: err.Error(),
			Data:    nil,
		})
	}

	users, err := h.userService.GetUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(400, domain.Response{
			Status:  "OK",
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(200, domain.Response{
		Status:  "OK",
		Message: "Get user successfully",
		Data:    users,
	})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	req := new(domain.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, domain.Response{
			Status:  "OK",
			Message: err.Error(),
			Data:    nil,
		})
	}

	err := h.userService.UpdateUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(400, domain.Response{
			Status:  "OK",
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(200, domain.Response{
		Status:  "OK",
		Message: "Update user successfully",
	})
}
