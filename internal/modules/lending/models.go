package lending

import (
	"time"

	"libro-system-api/internal/modules/books"
	"libro-system-api/internal/modules/users"

	"gorm.io/gorm"
)

type BookLending struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	BookInfoID  string         `gorm:"not null" json:"book_info_id"`                               // 關聯 BookInfo
	UserID      string         `gorm:"not null" json:"user_id"`                                    // 關聯 User
	LendingDate time.Time      `gorm:"not null" json:"lending_date"`                               // 借出日期
	ReturnDate  *time.Time     `json:"return_date"`                                                // 歸還日期（可為 null）
	DueDate     time.Time      `gorm:"not null" json:"due_date"`                                   // 歸還期限
	Status      string         `gorm:"type:varchar(20);not null;default:'borrowed'" json:"status"` // 狀態（borrowed, returned）
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"` // 軟刪除

	// 關聯
	BookInfo books.BookInfo `gorm:"foreignKey:BookInfoID;references:ID"`
	User     users.User     `gorm:"foreignKey:UserID;references:ID"`
}
