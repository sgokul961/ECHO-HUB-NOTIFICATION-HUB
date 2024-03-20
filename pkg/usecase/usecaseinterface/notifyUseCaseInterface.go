package NotifyusecaseinterfaceU

import (
	"github.com/sgokul961/echo-hub-notification-svc/pkg/models"
)

type NotifyUsecaseInterface interface {
	AddLikeNotification(notification models.LikeNotification) (int64, error)
}
