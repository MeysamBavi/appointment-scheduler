package main

import (
	"fmt"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/config"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/postgres"
	"log"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()
	db, err := postgres.Connect(cfg.Postgres)
	if err != nil {
		log.Fatal("could not connect to postgres", err)
	}
	service := app.NewHTTPService(
		app.Config{
			Port: 8080,
			CORS: cfg.CORS,
		},
		jwt.NewJWT("sth"),
		db,
	)

	service.Start()
}

func connectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(fmt.Sprintf("failed in connecting to database: %+v", err))
	}
	return db
}
