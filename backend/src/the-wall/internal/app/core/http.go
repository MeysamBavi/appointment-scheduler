package core

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Config struct {
	Port int
}

type HTTPService struct {
	server *echo.Echo
	config Config
}

func NewHTTPService(config Config) *HTTPService {
	e := echo.New()
	initRoutes(e)

	return &HTTPService{
		server: e,
		config: config,
	}
}

func initRoutes(e *echo.Echo) {
	e.GET("/login", login)
}

func (s *HTTPService) Start() {
	s.server.Logger.Fatal(s.server.Start(fmt.Sprintf(":%d", s.config.Port)))
}
