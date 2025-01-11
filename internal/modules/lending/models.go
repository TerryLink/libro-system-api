package lending

import (
	"time"

	"libro-system-api/internal/modules/books"
	"libro-system-api/internal/modules/users"

	"gorm.io/gorm"
)

type BookLending struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	BookInfoID  string         `gorm:"not null" json:"book_info_id"`
	UserID      string         `gorm:"not null" json:"user_id"`
	LendingDate time.Time      `gorm:"not null" json:"lending_date"`
	ReturnDate  *time.Time     `json:"return_date"`
	DueDate     time.Time      `gorm:"not null" json:"due_date"`
	Status      string         `gorm:"type:varchar(20);not null;default:'borrowed'" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// 關聯
	BookInfo books.BookInfo `gorm:"foreignKey:BookInfoID;references:ID"`
	User     users.User     `gorm:"foreignKey:UserID;references:ID"`
}
