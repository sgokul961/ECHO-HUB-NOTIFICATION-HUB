package domain

import "time"

type NotificationType string

const (
	CommentNotification NotificationType = "comment"
	LikeNotification    NotificationType = "like"
	FollowNotification  NotificationType = "follow"
	// Add more notification types as needed
)

// Notification represents the notification table in the database
type Notification struct {
	ID               int64            `json:"id" gorm:"primaryKey"`
	RecipientID      int64            `json:"recipient_id"`
	SenderID         int64            `json:"sender_id"`
	NotificationType NotificationType `json:"notification_type"`
	PostID           int64            `json:"post_id"`
	CommentID        int64            `json:"comment_id"`
	Seen             bool             `json:"seen" gorm:"default:false"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}
