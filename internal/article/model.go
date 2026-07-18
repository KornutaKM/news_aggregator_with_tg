package article

import "time"

type Article struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Link        string    `json:"link" gorm:"not null;unique"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	TopicID     uint      `json:"topic_id" gorm:"index;not null"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`

	Sentiment string   `json:"sentiment,omitempty"`
	Keywords  []string `json:"keywords,omitempty" gorm:"-"`
}
