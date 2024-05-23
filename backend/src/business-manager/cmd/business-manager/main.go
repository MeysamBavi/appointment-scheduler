package main

import (
	"fmt"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	service := app.NewHTTPService(
		app.Config{
			Port:       8080,
			EnableCORS: true,
		},
		jwt.NewJWT("sth"),
		connectDatabase(),
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
