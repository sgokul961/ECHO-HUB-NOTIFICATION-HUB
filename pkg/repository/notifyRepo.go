package repository

import (
	"github.com/sgokul961/echo-hub-notification-svc/pkg/domain"
	interfaceR "github.com/sgokul961/echo-hub-notification-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type NotificationDataBase struct {
	DB *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) interfaceR.NotifyRepoInterface {
	return &NotificationDataBase{
		DB: db,
	}
}
func (r *NotificationDataBase) AddNotification(notification domain.Notification) (int64, error) {

	err := r.DB.Create(&notification).Error
	if err != nil {
		// Handle error
		return 0, err
	}
	return notification.ID, nil
}
