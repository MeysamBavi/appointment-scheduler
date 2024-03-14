package notification

import "fmt"

type consoleLogger struct{}

func NewConsoleLogger() Notificator {
	return &consoleLogger{}
}

func (f *consoleLogger) SendNotification(phoneNumber string, content string) error {
	fmt.Printf("Hello %s. %s\n", phoneNumber, content)

	return nil
}
