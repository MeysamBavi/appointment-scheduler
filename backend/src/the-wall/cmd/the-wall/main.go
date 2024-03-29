package main

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/kvstore"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/notification"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/app/core"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/jwt"
)

func main() {
	service := core.NewHTTPService(
		core.Config{
			Port:       8080,
			EnableCORS: true,
		},
		kvstore.NewMemoryKVStore(),
		notification.NewConsoleLogger(),
		jwt.NewJWT("sth"),
	)

	service.Start()
}
