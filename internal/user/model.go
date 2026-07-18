package user

import "time"

type User struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"-" gorm:"unique;not null"`
	CreateAt time.Time `json:"create_at"`
}
