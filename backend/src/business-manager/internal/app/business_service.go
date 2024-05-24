package app

import (
	"errors"
	"net/http"

	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/handlers"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"github.com/labstack/echo/v4"
)

type createBusinessServiceRequest struct {
	Business uint   `param:"business_id"`
	Name     string `json:"name"`
}

func (s *HTTPService) CreateBusinessService(ctx echo.Context) error {
	request := createBusinessServiceRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{internalError})
	}
	isOwner, err := checkUserIsBusinessOwner(ctx, s.db, request.Business)
	if err != nil {
		response, status := handleBusinessOwnerPermissionError(err)
		return ctx.JSON(status, response)
	}
	if !isOwner {
		return ctx.JSON(http.StatusForbidden, &MessageResponse{Message: "you aren't business owner."})
	}

	if err = handlers.CreateBusinessService(s.db, &models.BusinessService{
		Name:       request.Name,
		BusinessID: request.Business,
	}); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{internalError})
	}

	return ctx.JSON(http.StatusCreated, &MessageResponse{"service created."})
}

type getBusinessServicesRequest struct {
	Business uint `param:"business_id"`
}

type getBusinessServicesResponse struct {
	Message          string                   `json:"message"`
	BusinessServices []models.BusinessService `json:"business_services"`
}

func (s *HTTPService) GetBusinessServices(ctx echo.Context) error {
	request := getBusinessServicesRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessServicesResponse{Message: internalError})
	}

	businessServices, err := handlers.GetBusinessServices(s.db, request.Business)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessServicesResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &getBusinessServicesResponse{
		BusinessServices: businessServices,
		Message:          "services retrieved.",
	})
}

type getBusinessServiceRequest struct {
	BusinessService uint `param:"service_id"`
	Business        uint `param:"business_id"`
}

type getBusinessServiceResponse struct {
	Message         string                  `json:"message"`
	BusinessService *models.BusinessService `json:"business_service,omitempty"`
}

func (s *HTTPService) GetBusinessService(ctx echo.Context) error {
	request := getBusinessServiceRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessServiceResponse{Message: internalError})
	}

	businessService, err := handlers.GetBusinessService(s.db, request.BusinessService, request.Business)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &getBusinessServiceResponse{Message: "service not found."})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessServiceResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &getBusinessServiceResponse{
		BusinessService: businessService,
		Message:         "service retrieved.",
	})
}

type deleteBusinessServiceRequest struct {
	BusinessService uint `param:"service_id"`
	Business        uint `param:"business_id"`
}

func (s *HTTPService) DeleteBusinessService(ctx echo.Context) error {
	request := deleteBusinessServiceRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}
	isOwner, err := checkUserIsBusinessOwner(ctx, s.db, request.Business)
	if err != nil {
		response, status := handleBusinessOwnerPermissionError(err)
		return ctx.JSON(status, response)
	}
	if !isOwner {
		return ctx.JSON(http.StatusForbidden, &MessageResponse{Message: "you aren't business owner."})
	}

	if err = handlers.DeleteBusinessService(s.db, request.BusinessService, request.Business); err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			ctx.Logger().Error(err)
			return ctx.JSON(http.StatusNotFound, &MessageResponse{Message: "service not found."})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &MessageResponse{
		Message: "service deleted.",
	})
}

type updateBusinessServiceRequest struct {
	BusinessService uint `param:"service_id"`
	createBusinessServiceRequest
}

func (s *HTTPService) UpdateBusinessService(ctx echo.Context) error {
	request := updateBusinessServiceRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}
	isOwner, err := checkUserIsBusinessOwner(ctx, s.db, request.Business)
	if err != nil {
		response, status := handleBusinessOwnerPermissionError(err)
		return ctx.JSON(status, response)
	}
	if !isOwner {
		return ctx.JSON(http.StatusForbidden, &MessageResponse{Message: "you aren't business owner."})
	}

	if err = handlers.UpdateBusinessService(s.db, request.BusinessService, &models.BusinessService{
		Name: request.Name,
	}); err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &MessageResponse{Message: "service not found"})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &MessageResponse{
		Message: "service updated.",
	})
}
