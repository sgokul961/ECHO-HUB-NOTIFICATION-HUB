package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/models"
)

func StartLikeKafkaConsumer() {
	brokerUrl := []string{"localhost:9092"}
	// Use the same Kafka broker URL as the producer

	topic := "like_notification"

	worker, err := ConnectToLikeConsumer(brokerUrl)
	if err != nil {
		panic(err)
	}

	partitionConsumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("consumer started")

	signchan := make(chan os.Signal, 1)
	signal.Notify(signchan, syscall.SIGINT, syscall.SIGTERM)
	messagecount := 0

	go func() {
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

				messagecount++
				fmt.Printf("received message count: %d | Topic: %s | Message: %s\n", messagecount, string(msg.Topic), string(msg.Value))

			case <-signchan:
				fmt.Println("interruption detected")
				donech <- struct{}{}
			}
		}
	}()

	<-donech

	fmt.Println("processed", messagecount, "messages")
	if err := worker.Close(); err != nil {
		panic(err)
	}
}

var donech = make(chan struct{})

func ConnectToLikeConsumer(brokerUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokerUrl, config)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
