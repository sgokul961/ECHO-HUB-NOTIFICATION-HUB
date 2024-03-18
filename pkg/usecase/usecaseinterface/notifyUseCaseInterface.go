package NotifyusecaseinterfaceU

import "github.com/sgokul961/echo-hub-notification-svc/pkg/domain"

type NotifyUsecaseInterface interface {
	AddCommentNotification(notification domain.Notification) (int64, error)
}
