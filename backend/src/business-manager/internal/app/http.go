package app

import (
	"fmt"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Config struct {
	Port int
	CORS httpserver.CORSConfig
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
	e.Use(httpserver.CORSMiddleware(config.CORS))
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

	e.POST("/businesses/:business_id/employees", service.CreateEmployee, httpserver.JWTMiddleware(service.jwtSdk))
	e.GET("/businesses/:business_id/employees", service.GetEmployees, httpserver.JWTMiddleware(service.jwtSdk))
	e.GET(
		"/businesses/:business_id/employees/:employee_id",
		service.GetEmployee,
		httpserver.JWTMiddleware(service.jwtSdk),
	)
	e.DELETE(
		"/businesses/:business_id/employees/:employee_id",
		service.DeleteEmployee,
		httpserver.JWTMiddleware(service.jwtSdk),
	)

	e.POST(
		"/businesses/:business_id/services",
		service.CreateBusinessService,
		httpserver.JWTMiddleware(service.jwtSdk),
	)
	e.GET("/businesses/:business_id/services", service.GetBusinessServices)
	e.GET("/businesses/:business_id/services/:service_id", service.GetBusinessService)
	e.DELETE(
		"/businesses/:business_id/services/:service_id",
		service.DeleteBusinessService,
		httpserver.JWTMiddleware(service.jwtSdk),
	)
	e.PUT(
		"/businesses/:business_id/services/:service_id",
		service.UpdateBusinessService,
		httpserver.JWTMiddleware(service.jwtSdk),
	)

	e.POST("/businesses", service.CreateBusiness, httpserver.JWTMiddleware(service.jwtSdk))
	e.GET("/businesses", service.GetBusinesses, httpserver.JWTMiddleware(service.jwtSdk))
	e.GET("/businesses/:business_id", service.GetBusiness, httpserver.JWTMiddleware(service.jwtSdk))
	e.PUT("/businesses/:business_id", service.UpdateBusiness, httpserver.JWTMiddleware(service.jwtSdk))
	e.DELETE("/businesses/:business_id", service.DeleteBusiness, httpserver.JWTMiddleware(service.jwtSdk))
}

func migrateDatabase(db *gorm.DB) {
	for _, model := range []any{
		&models.ServiceType{},
		&models.Business{},
		&models.Employee{},
		&models.BusinessService{},
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
			Name: "پزشکی",
		})
		db.Create(&models.ServiceType{
			Name: "آرایشی",
		})
	}
}

func (s *HTTPService) Start() {
	s.server.Logger.Fatal(s.server.Start(fmt.Sprintf(":%d", s.config.Port)))
}
