package httpserver

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	userIdContextKey = "userId"
)

func GetRequestToken(c echo.Context) (string, bool) {
	values, tokenPresent := c.Request().Header["Authorization"]
	if !tokenPresent || len(values) == 0 {
		return "", false
	}
	token := strings.TrimPrefix(values[0], "Bearer ")
	return token, token != ""
}

func JWTMiddleware(jwt *jwt.JWT) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, ok := GetRequestToken(c)
			if !ok {
				return echo.ErrUnauthorized
			}
			payload, err := jwt.ExtractPayload(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "malformed token", err)
			}

			c.Set(userIdContextKey, payload.UserId)
			return next(c)
		}
	}
}

func GetUserId(c echo.Context) (int32, bool) {
	val := c.Get(userIdContextKey)
	if val == nil {
		return 0, false
	}
	id, ok := val.(int32)
	return id, ok
}
