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

// func (n *NotificationHandler) ConsumeKafkaMessages(empty *pb.Empty, stream pb.NotificationService_ConsumeKafkaMessagesServer) error {
// 	ctx := stream.Context()

// 	// Consume Kafka messages using the NotifyUseCase method
// 	// Consume Kafka messages using the NotifyUseCase method
// 	for {
// 		kafkaMessage, err := n.notifyUsecase.ConsumeMessage()
// 		if err != nil {
// 			// Check if context is done or canceled
// 			select {
// 			case <-ctx.Done():
// 				return ctx.Err()
// 			default:
// 				// Handle other errors
// 				return err
// 			}
// 		}
// 		// Create a NotificationMessage instance from the Kafka message
// 		notification := &pb.NotificationMessage{
// 			PostId:  kafkaMessage.PostID,
// 			UserId:  kafkaMessage.UserID,
// 			Message: kafkaMessage.Message,
// 		}

// 		// Send the NotificationMessage to the client stream
// 		if err := stream.Send(notification); err != nil {
// 			// Handle error
// 			return err
// 		}
// 	}
// }

func (n *NotificationHandler) ConsumeKafkaMessages(empty *pb.Empty, stream pb.NotificationService_ConsumeKafkaMessagesServer) error {
	// Consume Kafka messages using the NotifyUseCase method
	for {
		kafkaMessage, err := n.notifyUsecase.ConsumeMessage()
		if err != nil {
			// Handle error
			return err
		}

		// Create a NotificationMessage instance from the Kafka message
		notification := &pb.NotificationMessage{
			PostId:  kafkaMessage.PostID,
			UserId:  kafkaMessage.UserID,
			Message: kafkaMessage.Message,
		}

		//Send the NotificationMessage to the client stream
		if err := stream.Send(notification); err != nil {
			// Handle error
			return err
		}

	}

}
