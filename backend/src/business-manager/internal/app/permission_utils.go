package app

import (
	"errors"
	"net/http"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	errUnauthorized     = errors.New("you are not authorized")
	errBusinessNotFound = errors.New("business not found")
)

func checkUserIsBusinessOwner(ctx echo.Context, db *gorm.DB, businessID uint) (bool, error) {
	userID, ok := httpserver.GetUserId(ctx)
	if !ok {
		return false, errUnauthorized
	}

	business, err := handlers.GetBusiness(db, businessID)
	if err != nil {
		if errors.Is(err, handlers.ErrNoRows) {
			return false, errBusinessNotFound
		}

		return false, err
	}

	return business.UserID == userID, nil
}

func handleBusinessOwnerPermissionError(err error) (string, int) {
	var status int
	var response string
	if errors.Is(err, errUnauthorized) {
		response = "you are not authorized."
		status = http.StatusUnauthorized
	} else if errors.Is(err, errBusinessNotFound) {
		response = "business not found."
		status = http.StatusNotFound
	} else {
		response = internalError
		status = http.StatusInternalServerError
	}
	return response, status
}
