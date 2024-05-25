package main

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/kvstore"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/notification"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/config"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/postgres"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/app/core"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/repo"
	"log"
)

func main() {
	cfg := config.Load()
	db, err := postgres.Connect(cfg.Postgres)
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
	}
	userRepo, err := repo.NewUser(db)
	if err != nil {
		log.Fatal("failed to connect to user database", err)
	}

	service := core.NewHTTPService(
		core.Config{
			Port: 8080,
			CORS: cfg.CORS,
		},
		kvstore.NewMemoryKVStore(),
		notification.NewConsoleLogger(),
		jwt.NewJWT("sth"),
		userRepo,
	)

	service.Start()
}
