package api

import (
	"log"
	"net"

	"github.com/sgokul961/echo-hub-notification-svc/pkg/api/handler"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/config"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *grpc.Server
}

func NewServerHttp(handler *handler.NotificationHandler) *ServerHTTP {
	engine := grpc.NewServer()
	pb.RegisterNotificationServiceServer(engine, handler)
	return &ServerHTTP{
		engine: engine,
	}

}
func (s *ServerHTTP) Start(c config.Config) {
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("error loading server", err)
	}

	if err = s.engine.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
