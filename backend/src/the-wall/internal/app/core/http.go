package core

import (
	"fmt"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/kvstore"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/notification"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/clients"
	"github.com/labstack/echo/v4"
)

type Config struct {
	Port int
}

type HTTPService struct {
	server *echo.Echo
	config Config

	otpClient clients.OTP
}

func NewHTTPService(
	config Config,
	kvStore kvstore.KVStore,
	notificator notification.Notificator,
) *HTTPService {
	e := echo.New()
	service := &HTTPService{
		server: e,
		config: config,

		otpClient: clients.NewOTPClient(kvStore, notificator),
	}

	initRoutes(e, service)

	return service
}

func initRoutes(e *echo.Echo, service *HTTPService) {
	e.POST("/otp/send", service.sendOTP)
	e.POST("/otp/validate", service.validateOTP)
}

func (s *HTTPService) Start() {
	s.server.Logger.Fatal(s.server.Start(fmt.Sprintf(":%d", s.config.Port)))
}
