package main

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/app/core"
)

func main() {
	service := core.NewHTTPService(core.Config{
		Port: 8080,
	})

	service.Start()
}
