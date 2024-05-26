package main

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"github.com/labstack/echo/v4"
	"github.com/swaggest/swgui/v5emb"
)

func main() {
	e := echo.New()
	cors := httpserver.CORSMiddleware(httpserver.CORSConfig{
		Enable: true,
	})
	docHandler := v5emb.New(
		"Appointment Scheduler",
		"/docs/openapi.json",
		"/docs/",
	)
	e.Group(
		"/docs/*",
		cors,
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return echo.WrapHandler(docHandler)
		},
	)
	e.File("/docs/openapi.json", "./openapi.json", cors)
	e.Logger.Fatal(e.Start(":8080"))
}
