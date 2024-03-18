package handler

import (
	"github.com/sgokul961/echo-hub-notification-svc/pkg/pb"
	NotifyusecaseinterfaceU "github.com/sgokul961/echo-hub-notification-svc/pkg/usecase/usecaseinterface"
)

type NotificationHandler struct {
	notifyUsecase NotifyusecaseinterfaceU.NotifyUsecaseInterface
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationHandler(notifyUseCase NotifyusecaseinterfaceU.NotifyUsecaseInterface) *NotificationHandler {
	return &NotificationHandler{
		notifyUsecase: notifyUseCase,
	}
}
