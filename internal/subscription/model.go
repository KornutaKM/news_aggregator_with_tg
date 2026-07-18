package subscription

import "time"

type Subscription struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	TopicID   uint      `json:"topic_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}
