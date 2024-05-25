package core

import (
	"fmt"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/repo"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/kvstore"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/notification"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/clients"
	"github.com/labstack/echo/v4"
)

type Config struct {
	Port int
	CORS httpserver.CORSConfig
}

type HTTPService struct {
	server *echo.Echo
	config Config

	otpClient clients.OTP
	jwtSdk    *jwt.JWT
	userRepo  repo.User
}

func NewHTTPService(
	config Config,
	kvStore kvstore.KVStore,
	notificator notification.Notificator,
	jwtSdk *jwt.JWT,
	userRepo repo.User,
) *HTTPService {
	e := echo.New()
	e.Use(httpserver.CORSMiddleware(config.CORS))
	service := &HTTPService{
		server: e,
		config: config,

		otpClient: clients.NewOTPClient(kvStore, notificator),
		jwtSdk:    jwtSdk,
		userRepo:  userRepo,
	}

	initRoutes(e, service)

	return service
}

func initRoutes(e *echo.Echo, service *HTTPService) {
	e.POST("/otp/send", service.sendOTP)
	e.POST("/otp/validate", service.validateOTP)
	e.GET("/authenticate", service.authenticateRequest)
	e.POST("/test", service.test, httpserver.JWTMiddleware(service.jwtSdk))
	e.GET("/users", service.getUsers)
}

func (s *HTTPService) Start() {
	s.server.Logger.Fatal(s.server.Start(fmt.Sprintf(":%d", s.config.Port)))
}
