package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSConfig struct {
	Enable  bool     `config:"enable"`
	Origins []string `config:"origins"`
}

func CORSMiddleware(config CORSConfig) echo.MiddlewareFunc {
	if !config.Enable {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return next(c)
			}
		}
	}
	echoConfig := middleware.DefaultCORSConfig
	if len(config.Origins) > 0 {
		echoConfig.AllowOrigins = config.Origins
	}
	return middleware.CORSWithConfig(echoConfig)
}
