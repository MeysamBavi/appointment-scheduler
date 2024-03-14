package notification

import (
	"strings"
)

func NormalizePhoneNumber(phoneNumber string) string {
	if strings.HasPrefix(phoneNumber, "+98") {
		return strings.Replace(phoneNumber, "+98", "0", 1)
	}

	if strings.HasPrefix(phoneNumber, "9") {
		return "0" + phoneNumber
	}

	return phoneNumber
}
