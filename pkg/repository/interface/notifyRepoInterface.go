package interfaceR

import "github.com/sgokul961/echo-hub-notification-svc/pkg/domain"

type NotifyRepoInterface interface {
	AddNotification(notification domain.Notification) (int64, error)
}
