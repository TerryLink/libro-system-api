package notification

import (
	"libro-system-api/internal/modules/users"
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	UserID    string         `gorm:"not null" json:"user_id"`
	Message   string         `gorm:"type:text;not null" json:"message"`
	SentAt    time.Time      `json:"sent_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	User users.User `gorm:"foreignKey:UserID;references:ID"`
}
