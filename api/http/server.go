package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"hrm/api/http/handler/auth"
	"hrm/api/http/handler/user"
	"hrm/config"
	"hrm/pkg/logger"
)

type ServInterface interface {
	Start(e *echo.Echo)
}

type Server struct {
	conf        *config.SystemConfig
	authHandler *auth.AuthHandler
	userHandler *user.UserHandler
}

func NewHttpServe(
	config *config.SystemConfig,
	authHandler *auth.AuthHandler,
	userHandler *user.UserHandler,
) *Server {
	return &Server{
		conf:        config,
		authHandler: authHandler,
		userHandler: userHandler,
	}
}

func (app *Server) Start(e *echo.Echo) {
	log := logger.GetLogger()
	err := app.InitRouters(e)
	if err != nil {
		log.Fatal().Msgf("InitRouters fail! %s", err)
	}

	httpPort := app.conf.HttpPort
	go func() {
		err := e.Start(fmt.Sprintf(":%d", httpPort))
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("can't start echo")
		}
	}()
	log.Info().Msg("all service already")
}
