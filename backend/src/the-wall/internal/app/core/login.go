package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type helloResponse struct {
	Message string `json:"message"`
}

func login(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &helloResponse{Message: "helllllllllo"})
}
