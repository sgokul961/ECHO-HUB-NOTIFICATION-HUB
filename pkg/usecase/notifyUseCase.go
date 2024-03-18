package usecase

import (
	"github.com/sgokul961/echo-hub-notification-svc/pkg/domain"
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
func (u *NotifyUseCase) AddCommentNotification(notification domain.Notification) (int64, error) {

	return 0, nil

}
