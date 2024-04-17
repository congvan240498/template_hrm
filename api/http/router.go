package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	prefixPath = "/api/v1"
)

func (app *Server) InitRouters(e *echo.Echo) error {
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	userClient := e.Group(prefixPath + "/user")
	userClient.POST("", app.userHandler.CreateUser)
	userClient.GET("", app.userHandler.GetUser)
	userClient.PUT("", app.userHandler.UpdateUser)

	e.POST(prefixPath+"/login", app.authHandler.Login)
	e.POST(prefixPath+"/logout", app.authHandler.Logout)

	return nil
}
