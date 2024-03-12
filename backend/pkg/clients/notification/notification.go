package notification

type Notificator interface {
	SendNotification(phoneNumber string, content string) error
}
