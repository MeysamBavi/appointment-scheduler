package app

import (
	"net/http"

	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/handlers"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"github.com/labstack/echo/v4"
)

type getServiceTypeRequest struct {
	Q string `query:"q"`
}

type getServiceTypeResponse struct {
	ServiceTypes []models.ServiceType `json:"service_types"`
	Message      string               `json:"message"`
}

func (s *HTTPService) GetServiceType(ctx echo.Context) error {
	request := getServiceTypeRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getServiceTypeResponse{Message: "internal error"})
	}

	serviceTypes, err := handlers.GetServiceTypes(s.db, request.Q)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getServiceTypeResponse{Message: "internal error"})
	}

	return ctx.JSON(http.StatusOK, getServiceTypeResponse{ServiceTypes: serviceTypes})
}
