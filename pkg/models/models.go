package models

type LikeNotification struct {
	ID      int64  `json:"id" gorm:"primaryKey"`
	UserID  int64  `json:"user_id"`
	Message string `json:"message"`
	PostID  int64  `json:"post_id"`
}
