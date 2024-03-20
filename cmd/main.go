package main

import (
	"log"

	"github.com/sgokul961/echo-hub-notification-svc/pkg/config"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/helper"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/wire"
)

func main() {
	// Load configuration
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed to load config:", err)
	}
	//calling kafka consumer

	go helper.StartKafkaConsumer()

	//go helper.StartLikeKafkaConsumer()

	// Initialize API server
	server, err := wire.InitApi(c)
	if err != nil {
		log.Fatalln("error initializing server:", err)
	}

	server.Start(c)

}
