package app

import (
	"errors"
	"net/http"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/handlers"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"github.com/labstack/echo/v4"
)

type createBusinessRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	ServiceType uint   `json:"service_type"`
}

func (s *HTTPService) CreateBusiness(ctx echo.Context) error {
	request := createBusinessRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{internalError})
	}

	if _, err = handlers.GetServiceType(s.db, request.ServiceType); err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &MessageResponse{"service type not found."})
		}

		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{internalError})
	}

	err = handlers.CreateBusiness(s.db, &models.Business{
		Name:          request.Name,
		Address:       request.Address,
		ServiceTypeID: request.ServiceType,
	})
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{internalError})
	}

	return ctx.JSON(http.StatusCreated, &MessageResponse{"business created."})
}

type getBusinessesResponse struct {
	Businesses []models.Business `json:"businesses"`
	Message    string            `json:"message"`
}

func (s *HTTPService) GetBusinesses(ctx echo.Context) error {
	userID, ok := httpserver.GetUserId(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, &getBusinessesResponse{
			Message: "you are not authorized.",
		})
	}
	businesses, err := handlers.GetBusinesses(s.db, userID)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessesResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusCreated, &getBusinessesResponse{
		Businesses: businesses,
		Message:    "businesses retrieved.",
	})
}

type getBusinessRequest struct {
	BusinessID uint `param:"business_id"`
}

type getBusinessResponse struct {
	Business *models.Business `json:"business,omitempty"`
	Message  string           `json:"message"`
}

func (s *HTTPService) GetBusiness(ctx echo.Context) error {
	request := getBusinessRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessResponse{Message: internalError})
	}

	business, err := handlers.GetBusiness(s.db, request.BusinessID)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &getBusinessResponse{Message: "business not found"})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &getBusinessResponse{
		Business: business,
		Message:  "business retrieved.",
	})
}

type updateBusinessRequest struct {
	BusinessID uint `param:"business_id"`
	createBusinessRequest
}

func (s *HTTPService) UpdateBusiness(ctx echo.Context) error {
	request := updateBusinessRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}
	isOwner, err := checkUserIsBusinessOwner(ctx, s.db, request.BusinessID)
	if err != nil {
		response, status := handleBusinessOwnerPermissionError(err)
		return ctx.JSON(status, response)
	}
	if !isOwner {
		return ctx.JSON(http.StatusForbidden, &MessageResponse{Message: "you aren't business owner."})
	}

	if err = handlers.UpdateBusiness(s.db, request.BusinessID, &models.Business{
		Name:          request.Name,
		Address:       request.Address,
		ServiceTypeID: request.ServiceType,
	}); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &MessageResponse{
		Message: "business updated.",
	})
}

type deleteBusinessRequest struct {
	BusinessID uint `param:"business_id"`
}

func (s *HTTPService) DeleteBusiness(ctx echo.Context) error {
	request := deleteBusinessRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}
	isOwner, err := checkUserIsBusinessOwner(ctx, s.db, request.BusinessID)
	if err != nil {
		response, status := handleBusinessOwnerPermissionError(err)
		return ctx.JSON(status, response)
	}
	if !isOwner {
		return ctx.JSON(http.StatusForbidden, &MessageResponse{Message: "you aren't business owner."})
	}

	if err = handlers.DeleteBusiness(s.db, request.BusinessID); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &MessageResponse{
		Message: "business deleted.",
	})
}
