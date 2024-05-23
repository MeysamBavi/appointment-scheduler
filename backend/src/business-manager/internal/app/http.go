package app

import (
	"fmt"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Config struct {
	Port       int
	EnableCORS bool
}

type HTTPService struct {
	server *echo.Echo
	config Config

	jwtSdk *jwt.JWT
	db     *gorm.DB
}

func NewHTTPService(
	config Config,
	jwtSdk *jwt.JWT,
	db *gorm.DB,
) *HTTPService {
	e := echo.New()
	if config.EnableCORS {
		e.Use(middleware.CORS())
	}
	service := &HTTPService{
		server: e,
		config: config,

		jwtSdk: jwtSdk,
		db:     db,
	}

	initRoutes(e, service)
	migrateDatabase(db)

	return service
}

func initRoutes(e *echo.Echo, service *HTTPService) {
	e.GET("/service_types", service.GetServiceType)
}

func migrateDatabase(db *gorm.DB) {
	for _, model := range []any{
		&models.ServiceType{},
	} {
		err := db.AutoMigrate(model)
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database %+v", err))
		}
	}

	// TODO: remove these rows for production
	var sampleServiceType models.ServiceType
	result := db.First(&sampleServiceType)
	if result.RowsAffected == 0 {
		db.Create(&models.ServiceType{
			Name: "service_type1",
		})
		db.Create(&models.ServiceType{
			Name: "servicetype2",
		})
	}
}

func (s *HTTPService) Start() {
	s.server.Logger.Fatal(s.server.Start(fmt.Sprintf(":%d", s.config.Port)))
}
