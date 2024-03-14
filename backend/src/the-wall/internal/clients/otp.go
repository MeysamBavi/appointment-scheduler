package clients

import (
	"fmt"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/kvstore"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/clients/notification"
)

type OTP interface {
	SendOTP(phoneNumber string) error
	ValidateOTP(phoneNumber string, code string) (bool, error)
}

func getKey(phoneNumber string) string {
	const otpKeyTemplate = "otp:%s" // e.g) otp:09904614116

	return fmt.Sprintf(otpKeyTemplate, phoneNumber)
}

type fakeOTP struct {
	store       kvstore.KVStore
	notificator notification.Notificator
}

func (o *fakeOTP) SendOTP(phoneNumber string) error {
	fakeCode := phoneNumber[11-4:]

	err := o.store.Set(getKey(phoneNumber), fakeCode)
	if err != nil {
		return err
	}

	err = o.notificator.SendNotification(phoneNumber, fakeCode)
	if err != nil {
		return err
	}

	return nil
}

func (o *fakeOTP) ValidateOTP(phoneNumber string, code string) (bool, error) {
	storedCode, err := o.store.Get(getKey(phoneNumber))
	if err != nil {
		return false, err
	}

	return storedCode == code, nil
}

func NewOTPClient(store kvstore.KVStore, notificator notification.Notificator) OTP {
	return &fakeOTP{
		store:       store,
		notificator: notificator,
	}
}
