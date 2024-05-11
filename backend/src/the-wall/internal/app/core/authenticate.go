package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	authorizationTokenHeaderKey = "Authorization"
)

func (s *HTTPService) authenticateRequest(ctx echo.Context) error {
	token, tokenPresent := ctx.Request().Header[authorizationTokenHeaderKey]
	if !tokenPresent || len(token) == 0 {
		return ctx.JSON(http.StatusUnauthorized, nil)
	}

	if err := s.jwtSdk.CheckValidity(token[0]); err != nil {
		return ctx.JSON(http.StatusUnauthorized, nil)
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (s *HTTPService) test(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
