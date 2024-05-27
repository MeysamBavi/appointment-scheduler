package core

import (
	"context"
	"errors"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/pkg/clients"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (s *HTTPService) getUsers(c echo.Context) error {
	idStr := c.QueryParam("id")
	if idStr != "" {
		idBig, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		id := uint(idBig)

		user, err := s.userRepo.Get(context.Background(), id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, &clients.User{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
		})
	}

	phone := c.QueryParam("phone")
	if phone == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	user, err := s.userRepo.GetByPhoneNumber(context.Background(), phone)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, &clients.User{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
	})
}
