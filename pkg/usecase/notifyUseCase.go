package usecase

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/helper"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/models"
	interfaceR "github.com/sgokul961/echo-hub-notification-svc/pkg/repository/interface"
	usecaseinterfaceU "github.com/sgokul961/echo-hub-notification-svc/pkg/usecase/usecaseinterface"
)

type NotifyUseCase struct {
	notifyRepo interfaceR.NotifyRepoInterface
}

func NewUserUserUseCase(repo interfaceR.NotifyRepoInterface) usecaseinterfaceU.NotifyUsecaseInterface {
	return &NotifyUseCase{
		notifyRepo: repo,
	}
}

var doneCh = make(chan struct{})

// func (u *NotifyUseCase) AddLikeNotification(notification models.LikeNotification) (int64, error) {

// 	var userID int64

// 	brokerUrl := []string{"localhost:9092"}
// 	// Use the same Kafka broker URL as the producer

// 	topic := "like_notification"

// 	worker, err := helper.ConnectToLikeConsumer(brokerUrl)
// 	if err != nil {
// 		panic(err)
// 	}

// 	partitionConsumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("consumer started")

// 	signchan := make(chan os.Signal, 1)
// 	signal.Notify(signchan, syscall.SIGINT, syscall.SIGTERM)
// 	messagecount := 0

// 	go func() {
// 		for {
// 			select {
// 			case err := <-partitionConsumer.Errors():
// 				fmt.Println("Error:", err)
// 			case msg := <-partitionConsumer.Messages():

// 				var message models.LikeNotification

// 				if err := json.Unmarshal(msg.Value, &message); err != nil {
// 					fmt.Println("Error parsing message:", err)
// 					continue
// 				}

// 				// Add the notification to the repository
// 				userID, err := u.notifyRepo.AddNotification(domain.LikeNotification{
// 					UserID:  message.UserID,
// 					PostID:  message.PostID,
// 					Message: message.Message,
// 				})
// 				if err != nil {
// 					fmt.Println("Error adding notification to repository:", err)
// 					continue
// 				}

// 				fmt.Printf("Notification added for UserID: %d\n", userID)

// 				messagecount++
// 				fmt.Printf("received message count: %d | Topic: %s | Message: %s\n", messagecount, string(msg.Topic), string(msg.Value))
// 			case <-signchan:
// 				fmt.Println("interruption detected")
// 				donech <- struct{}{}
// 			}
// 		}
// 	}()

// 	// Wait for the consumer to finish processing messages
// 	<-donech

//		fmt.Println("processed", messagecount, "messages")
//		if err := worker.Close(); err != nil {
//			panic(err)
//		}
//		return userID, nil
//	}
func (u *NotifyUseCase) AddLikeNotification(notification models.LikeNotification) (int64, error) {

	var userID int64

	var messagecount int

	brokerURL := []string{"localhost:9092"}

	topic := "like_notification"

	worker, err := helper.ConnectToLikeConsumer(brokerURL)

	if err != nil {
		return 0, fmt.Errorf("failed to connect to Kafka consumer: %w", err)
	}
	defer worker.Close()

	partitionConsumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)

	if err != nil {
		return 0, fmt.Errorf("failed to consume partition: %w", err)
	}

	fmt.Println("consumer started")

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer close(doneCh)

		for {
			select {
			case err := <-partitionConsumer.Errors():
				fmt.Println("Error:", err)
			case msg := <-partitionConsumer.Messages():
				var message models.LikeNotification

				if err := json.Unmarshal(msg.Value, &message); err != nil {
					fmt.Println("Error parsing message:", err)
					continue
				}

				// Add the notification to the repository
				fmt.Println("postid", message.PostID)

				userID, err := u.notifyRepo.AddNotification(domain.LikeNotification{
					UserID:  message.UserID,
					PostID:  message.PostID,
					Message: "you got one like",
				})
				if err != nil {
					fmt.Println("Error adding notification to repository:", err)
					continue
				}

				fmt.Printf("Notification added for UserID: %d\n", userID)

				messagecount++
				fmt.Printf("received message count: %d | Topic: %s | Message: %s\n", messagecount, string(msg.Topic), string(msg.Value))
			case <-signChan:
				fmt.Println("interruption detected")
				return
			}
		}
	}()

	<-doneCh
	fmt.Println("processed", messagecount, "messages")

	return userID, nil
}
