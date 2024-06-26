package core

import (
	"context"
	"errors"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/models"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"time"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/kvstore"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/notification"
	"github.com/labstack/echo/v4"
)

func isPhoneNumberValid(phoneNumber string) (bool, error) {
	reg, err := regexp.Compile(notification.PhoneNumberRegex)
	if err != nil {
		return false, err
	}

	if !reg.MatchString(phoneNumber) {
		return false, nil
	}

	return true, nil
}

type sendOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type sendOTPResponse struct {
	Message string `json:"message"`
}

var (
	sendOTPInternalErrorResponse        = sendOTPResponse{Message: "Some problem occurred in sending otp."}
	otpCodeSendResponse                 = sendOTPResponse{Message: "Otp code sent. Check your phone."}
	sendOTPPhoneNumberIsInvalidResponse = sendOTPResponse{Message: "Phone number is invalid."}
)

func (s *HTTPService) sendOTP(ctx echo.Context) error {
	request := sendOTPRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, sendOTPInternalErrorResponse)
	}

	phoneIsValid, err := isPhoneNumberValid(request.PhoneNumber)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, sendOTPInternalErrorResponse)
	}
	if !phoneIsValid {
		return ctx.JSON(http.StatusBadRequest, sendOTPPhoneNumberIsInvalidResponse)
	}

	err = s.otpClient.SendOTP(notification.NormalizePhoneNumber(request.PhoneNumber))
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, sendOTPInternalErrorResponse)
	}

	return ctx.JSON(http.StatusOK, otpCodeSendResponse)
}

type validateOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}

type validateOTPResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

var (
	validateInternalErrorResponse = validateOTPResponse{
		Message: "Some internal error occurred in validating otp.",
	}
	otpIsInvalidResponse                   = validateOTPResponse{Message: "Send a valid otp."}
	validteOTPPhoneNumberIsInvalidResponse = sendOTPResponse{Message: "Phone number is invalid."}
)

func (s *HTTPService) validateOTP(ctx echo.Context) error {
	request := validateOTPRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, validateInternalErrorResponse)
	}

	phoneIsValid, err := isPhoneNumberValid(request.PhoneNumber)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, validateInternalErrorResponse)
	}
	if !phoneIsValid {
		return ctx.JSON(http.StatusBadRequest, validteOTPPhoneNumberIsInvalidResponse)
	}

	isValid, err := s.otpClient.ValidateOTP(notification.NormalizePhoneNumber(request.PhoneNumber), request.Code)
	if err != nil {
		if errors.Is(err, kvstore.KeyDoesNotExist) {
			return ctx.JSON(http.StatusBadRequest, otpIsInvalidResponse)
		}
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, validateInternalErrorResponse)
	}

	if !isValid {
		return ctx.JSON(http.StatusBadRequest, otpIsInvalidResponse)
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := s.userRepo.GetByPhoneNumber(ctx2, request.PhoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, validateInternalErrorResponse)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u := models.User{
			PhoneNumber: request.PhoneNumber,
		}
		err = s.userRepo.Create(ctx2, &u)
		user = &u
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, validateInternalErrorResponse)
	}

	jwtToken, err := s.jwtSdk.GetSignedJWT(jwt.Payload{
		UserId:      user.ID,
		PhoneNumber: user.PhoneNumber,
	})
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, validateInternalErrorResponse)
	}

	return ctx.JSON(http.StatusOK, validateOTPResponse{
		Message: "You are logged in.",
		Token:   jwtToken,
	})
}
