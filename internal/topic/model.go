package topic

import "time"

type Topic struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
}
