package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/pkg/clients"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/handlers"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"github.com/labstack/echo/v4"
)

type createEmployeeRequest struct {
	Business uint `param:"business_id"`
	User     uint `json:"user"`
}

func (s *HTTPService) CreateEmployee(ctx echo.Context) error {
	request := createEmployeeRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{internalError})
	}
	if request.User == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			&MessageResponse{"you should send user id to create employee."},
		)
	}

	isOwner, err := checkUserIsBusinessOwner(ctx, s.db, request.Business)
	if err != nil {
		response, status := handleBusinessOwnerPermissionError(err)
		return ctx.JSON(status, &MessageResponse{Message: response})
	}
	if !isOwner {
		return ctx.JSON(http.StatusForbidden, &MessageResponse{Message: "you aren't business owner."})
	}

	if err = handlers.CreateEmployee(s.db, &models.Employee{
		UserID:     request.User,
		BusinessID: request.Business,
	}); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ctx.JSON(http.StatusConflict, &MessageResponse{"employee already exist."})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{internalError})
	}

	return ctx.JSON(http.StatusCreated, &MessageResponse{"employee created."})
}

type getEmployeesRequest struct {
	Business uint `param:"business_id"`
}

type getEmployeesResponse struct {
	Message   string             `json:"message"`
	Employees []employeeResponse `json:"employees"`
}

type employeeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func serializeGetEmployeesResponse(wallClient *clients.TheWall, employees []models.Employee) *getEmployeesResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serializedEmployees := make([]employeeResponse, len(employees))
	for i, employee := range employees {
		user, err := wallClient.GetUserById(ctx, employee.UserID)
		if err != nil {
			log.Println(err)
			continue
		}
		serializedEmployees[i] = employeeResponse{
			ID:   employee.ID,
			Name: strings.TrimSpace(fmt.Sprintf("%s %s", user.Firstname, user.Lastname)),
		}
	}
	return &getEmployeesResponse{
		Employees: serializedEmployees,
		Message:   "businesses retrieved.",
	}
}

func (s *HTTPService) GetEmployees(ctx echo.Context) error {
	request := getEmployeesRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeesResponse{Message: internalError})
	}
	isOwner, err := checkUserIsBusinessOwner(ctx, s.db, request.Business)
	if err != nil {
		response, status := handleBusinessOwnerPermissionError(err)
		return ctx.JSON(status, response)
	}
	if !isOwner {
		return ctx.JSON(http.StatusForbidden, &MessageResponse{Message: "you aren't business owner."})
	}

	employees, err := handlers.GetEmployees(s.db, request.Business)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeesResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, serializeGetEmployeesResponse(s.wallClient, employees))
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
		return ctx.JSON(http.StatusInternalServerError, &getEmployeeResponse{Message: internalError})
	}
	isOwner, err := checkUserIsBusinessOwner(ctx, s.db, request.Business)
	if err != nil {
		response, status := handleBusinessOwnerPermissionError(err)
		return ctx.JSON(status, response)
	}
	if !isOwner {
		return ctx.JSON(http.StatusForbidden, &MessageResponse{Message: "you aren't business owner."})
	}

	employee, err := handlers.GetEmployee(s.db, request.Employee, request.Business)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, &getEmployeeResponse{Message: "employee not found."})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &getEmployeeResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &getEmployeeResponse{Employee: employee, Message: "employees retrieved."})
}

type deleteEmployeeRequest struct {
	Employee uint `param:"employee_id"`
	Business uint `param:"business_id"`
}

func (s *HTTPService) DeleteEmployee(ctx echo.Context) error {
	request := deleteEmployeeRequest{}
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

	if err = handlers.DeleteEmployee(s.db, request.Employee, request.Business); err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			ctx.Logger().Error(err)
			return ctx.JSON(http.StatusNotFound, &MessageResponse{Message: "employee not found."})
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, &MessageResponse{Message: internalError})
	}

	return ctx.JSON(http.StatusOK, &MessageResponse{
		Message: "employee deleted.",
	})
}
