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

type createBusinessResponse struct {
	Message string `json:"message"`
}

func (s *HTTPService) CreateBusiness(ctx echo.Context) error {
	request := createBusinessRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &createBusinessResponse{"internal error"})
	}

	err = handlers.CreateBusiness(s.db, &models.Business{
		Name:          request.Name,
		Address:       request.Address,
		ServiceTypeID: request.ServiceType,
	})
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &createBusinessResponse{"internal error"})
	}

	return ctx.JSON(http.StatusCreated, &createBusinessResponse{"business created."})
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
	businesses, err := handlers.GetBusinesses(s.db, uint(userID))
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessesResponse{Message: "internal error"})
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
		return ctx.JSON(http.StatusInternalServerError, &getBusinessResponse{Message: "internal error"})
	}

	business, err := handlers.GetBusiness(s.db, request.BusinessID)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &getBusinessResponse{Message: "business not found"})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getBusinessResponse{Message: "internal error"})
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

type updateBusinessResponse struct {
	Message string `json:"message"`
}

func (s *HTTPService) UpdateBusiness(ctx echo.Context) error {
	request := updateBusinessRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &updateBusinessResponse{Message: "internal error"})
	}
	userID, ok := httpserver.GetUserId(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, &updateBusinessResponse{
			Message: "you are not authorized.",
		})
	}

	business, err := handlers.GetBusiness(s.db, request.BusinessID)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &updateBusinessResponse{Message: "business not found"})
		}

		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &updateBusinessResponse{Message: "internal error"})
	}
	if business.UserID != uint(userID) {
		return ctx.JSON(http.StatusForbidden, &updateBusinessResponse{Message: "you aren't business owner."})
	}

	if err = handlers.UpdateBusiness(s.db, request.BusinessID, &models.Business{
		Name:          request.Name,
		Address:       request.Address,
		ServiceTypeID: request.ServiceType,
	}); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &updateBusinessResponse{Message: "internal error"})
	}

	return ctx.JSON(http.StatusOK, &updateBusinessResponse{
		Message: "business updated.",
	})
}

type deleteBusinessRequest struct {
	BusinessID uint `param:"business_id"`
}

type deleteBusinessResponse struct {
	Message string `json:"message"`
}

func (s *HTTPService) DeleteBusiness(ctx echo.Context) error {
	request := deleteBusinessRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &deleteBusinessResponse{Message: "internal error"})
	}
	userID, ok := httpserver.GetUserId(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, &deleteBusinessResponse{
			Message: "you are not authorized.",
		})
	}

	business, err := handlers.GetBusiness(s.db, request.BusinessID)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &deleteBusinessResponse{Message: "business not found"})
		}

		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &deleteBusinessResponse{Message: "internal error"})
	}
	if business.UserID != uint(userID) {
		return ctx.JSON(http.StatusForbidden, &deleteBusinessResponse{Message: "you aren't business owner."})
	}

	if err = handlers.DeleteBusiness(s.db, request.BusinessID); err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &deleteBusinessResponse{Message: "internal error"})
	}

	return ctx.JSON(http.StatusOK, &deleteBusinessResponse{
		Message: "business deleted.",
	})
}
