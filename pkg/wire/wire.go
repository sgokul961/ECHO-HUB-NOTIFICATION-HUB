//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/api"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/api/handler"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/config"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/db"
	postclient "github.com/sgokul961/echo-hub-notification-svc/pkg/postClient"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/repository"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/usecase"
)

func InitApi(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.Init,
		repository.NewNotificationRepo,
		usecase.NewUserUserUseCase,
		handler.NewNotificationHandler,
		postclient.NewPostServiceClient,

		api.NewServerHttp)
	return &api.ServerHTTP{}, nil

}
