package core

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *HTTPService) authenticateRequest(ctx echo.Context) error {
	token, tokenPresent := httpserver.GetRequestToken(ctx)
	if !tokenPresent {
		return ctx.JSON(http.StatusUnauthorized, nil)
	}

	if err := s.jwtSdk.CheckValidity(token); err != nil {
		return ctx.JSON(http.StatusUnauthorized, nil)
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (s *HTTPService) test(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
