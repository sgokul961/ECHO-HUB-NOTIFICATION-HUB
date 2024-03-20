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
func (r *NotificationDataBase) AddNotification(notification domain.LikeNotification) (int64, error) {
	// Write the SQL query using a prepared statement
	query := `INSERT INTO like_notifications (user_id, message, post_id) VALUES (?, ?, ?)`

	// Execute the SQL query with the provided values
	err := r.DB.Exec(query, notification.UserID, notification.Message, notification.PostID).Error
	if err != nil {
		return 0, err
	}

	return notification.UserID, nil
}
