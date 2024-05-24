package app

import (
	"errors"
	"net/http"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/handlers"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type createEmployeeRequest struct {
	Business uint `param:"business_id"`
	User     uint `json:"user"`
}

type createEmployeeResponse struct {
	Message string `json:"message"`
}

func (s *HTTPService) CreateEmployee(ctx echo.Context) error {
	request := createEmployeeRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &createEmployeeResponse{"internal error"})
	}
	userID, ok := httpserver.GetUserId(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, &createEmployeeResponse{
			Message: "you are not authorized.",
		})
	}

	if request.User == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			&createEmployeeResponse{"you should send user id to create employee."},
		)
	}

	business, err := handlers.GetBusiness(s.db, request.Business)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &createEmployeeResponse{"business not found."})
		}

		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &createEmployeeResponse{"internal error"})
	}
	if business.UserID != uint(userID) {
		return ctx.JSON(http.StatusForbidden, &createEmployeeResponse{Message: "you aren't business owner."})
	}

	// TODO: check permission

	if err = handlers.CreateEmployee(s.db, &models.Employee{
		UserID:     request.User,
		BusinessID: request.Business,
	}); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ctx.JSON(http.StatusConflict, &createEmployeeResponse{"employee already exist."})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &createEmployeeResponse{"internal error"})
	}

	return ctx.JSON(http.StatusCreated, &createEmployeeResponse{"employee created."})
}

type getEmployeesRequest struct {
	Business uint `param:"business_id"`
}

type getEmployeesResponse struct {
	Message   string            `json:"message"`
	Employees []models.Employee `json:"employees"`
}

func (s *HTTPService) GetEmployees(ctx echo.Context) error {
	request := getEmployeesRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeesResponse{Message: "internal error"})
	}
	userID, ok := httpserver.GetUserId(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, &getEmployeesResponse{
			Message: "you are not authorized.",
		})
	}

	business, err := handlers.GetBusiness(s.db, request.Business)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &getEmployeesResponse{Message: "business not found."})
		}

		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeesResponse{Message: "internal error"})
	}
	if business.UserID != uint(userID) {
		return ctx.JSON(http.StatusForbidden, &getEmployeesResponse{Message: "you aren't business owner."})
	}

	employees, err := handlers.GetEmployees(s.db, request.Business)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeesResponse{Message: "internal error"})
	}

	return ctx.JSON(http.StatusOK, &getEmployeesResponse{Employees: employees, Message: "employees retrieved."})
}

type getEmployeeRequest struct {
	Employee uint `param:"employee_id"`
	Business uint `param:"business_id"`
}

type getEmployeeResponse struct {
	Message  string           `json:"message"`
	Employee *models.Employee `json:"employee,omitempty"`
}

func (s *HTTPService) GetEmployee(ctx echo.Context) error {
	request := getEmployeeRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeeResponse{Message: "internal error"})
	}
	userID, ok := httpserver.GetUserId(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, &getEmployeeResponse{
			Message: "you are not authorized.",
		})
	}

	business, err := handlers.GetBusiness(s.db, request.Business)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &getEmployeeResponse{Message: "business not found."})
		}

		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeeResponse{Message: "internal error"})
	}
	if business.UserID != uint(userID) {
		return ctx.JSON(http.StatusForbidden, &getEmployeeResponse{Message: "you aren't business owner."})
	}

	employee, err := handlers.GetEmployee(s.db, request.Employee, request.Business)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &getEmployeeResponse{Message: "employee not found."})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeeResponse{Message: "internal error"})
	}

	return ctx.JSON(http.StatusOK, &getEmployeeResponse{Employee: employee, Message: "employees retrieved."})
}

type deleteEmployeeRequest struct {
	Employee uint `param:"employee_id"`
	Business uint `param:"business_id"`
}

type deleteEmployeeResponse struct {
	Message string `json:"message"`
}

func (s *HTTPService) DeleteEmployee(ctx echo.Context) error {
	request := deleteEmployeeRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &deleteEmployeeResponse{Message: "internal error"})
	}
	userID, ok := httpserver.GetUserId(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, &deleteEmployeeResponse{
			Message: "you are not authorized.",
		})
	}

	business, err := handlers.GetBusiness(s.db, request.Business)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &deleteEmployeeResponse{Message: "business not found."})
		}

		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &deleteEmployeeResponse{Message: "internal error"})
	}
	if business.UserID != uint(userID) {
		return ctx.JSON(http.StatusForbidden, &deleteEmployeeResponse{Message: "you aren't business owner."})
	}

	if err = handlers.DeleteEmployee(s.db, request.Employee, request.Business); err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			ctx.Logger().Error(err)
			return ctx.JSON(http.StatusNotFound, &deleteEmployeeResponse{Message: "employee not found."})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &deleteEmployeeResponse{Message: "internal error"})
	}

	return ctx.JSON(http.StatusOK, &deleteEmployeeResponse{
		Message: "employee deleted.",
	})
}
