package helper

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

// func StartKafkaConsumer() {
// 	// Define Kafka consumer configuration
// 	config := sarama.NewConfig()
// 	fmt.Println("conf", config)
// 	config.Consumer.Return.Errors = true

// 	// Create a new Kafka consumer
// 	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)

// 	fmt.Println("consumer", consumer)
// 	if err != nil {
// 		log.Fatalf("Error creating Kafka consumer: %v", err)
// 	}
// 	defer func() {
// 		if err := consumer.Close(); err != nil {
// 			log.Fatalf("Error closing Kafka consumer: %v", err)
// 		}
// 	}()

// 	// Create a new context and add a signal handler to gracefully shut down the consumer
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	signals := make(chan os.Signal, 1)
// 	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

// 	go func() {
// 		<-signals
// 		cancel()
// 	}()

// 	// Consume messages from the Kafka topic
// 	topic := "comment_notifications"
// 	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
// 	if err != nil {
// 		log.Fatalf("Error consuming topic %s: %v", topic, err)
// 	}
// 	defer func() {
// 		if err := partitionConsumer.Close(); err != nil {
// 			log.Fatalf("Error closing partition consumer: %v", err)
// 		}
// 	}()

// 	for {
// 		select {
// 		case msg := <-partitionConsumer.Messages():
// 			// Process Kafka message
// 			fmt.Printf("Received message: %s\n", string(msg.Value))
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
//}

func StartKafkaConsumer() {

	topic := "comment_notifications"

	brokerUrl := []string{"localhost:9092"}
	// Use the same Kafka broker URL as the producer

	worker, err := connectConsumer(brokerUrl)
	if err != nil {
		panic(err)
	}

	// // Obtain the partition IDs for the given topic
	// partitions, err := worker.Partitions(topic)
	// if err != nil {
	// 	panic(err)
	// }

	partitionConsumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("consumer started")

	signchan := make(chan os.Signal, 1)
	signal.Notify(signchan, syscall.SIGINT, syscall.SIGTERM)
	messagecount := 0

	//errors := make(chan error)

	// // Consume messages from each partition
	// for _, partition := range partitions {
	// 	go consumePartition(worker, topic, partition, errors)
	// }
	// select {
	// case err := <-errors:
	// 	fmt.Println("Error:", err)
	// case <-signchan:
	// 	fmt.Println("Received interrupt signal")
	// }

	go func() {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				fmt.Println("Error:", err)
			case msg := <-partitionConsumer.Messages():
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

func connectConsumer(brokerUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokerUrl, config)

	if err != nil {
		return nil, err
	}
	return conn, nil
}

// func consumePartition(worker sarama.Consumer, topic string, partition int32, errors chan<- error) {
// 	// Set the offset to the earliest available offset
// 	partitionConsumer, err := worker.ConsumePartition(topic, partition, sarama.OffsetOldest)
// 	if err != nil {
// 		errors <- err
// 		return
// 	}

// 	// Consume messages from the partition
// 	for msg := range partitionConsumer.Messages() {
// 		fmt.Printf("Topic: %s, Partition: %d, Offset: %d, Key: %s, Value: %s\n",
// 			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
// 	}

// 	// Close the partition consumer
// 	if err := partitionConsumer.Close(); err != nil {
// 		errors <- err
// 	}
// }
