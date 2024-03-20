package handler

import (
	"context"

	"github.com/sgokul961/echo-hub-notification-svc/pkg/models"
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
func (n *NotificationHandler) SendLikeNotification(ctx context.Context, notify *pb.LikeNotification) (*pb.NotificationResponse, error) {
	_, err := n.notifyUsecase.AddLikeNotification(models.LikeNotification{
		UserID: notify.UserId,
		PostID: notify.PostId,
	})
	if err != nil {
		return nil, err
	}

	// If there's no error, construct and return the response
	response := &pb.NotificationResponse{
		Message: "successfully got all notifications",
	}

	return response, nil

}
