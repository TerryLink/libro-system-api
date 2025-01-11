package books

import (
	"time"

	"libro-system-api/internal/modules/users"

	"gorm.io/gorm"
)

type Book struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	Title      string     `gorm:"type:varchar(100);not null" json:"title"`
	Author     string     `gorm:"type:varchar(100);not null" json:"author"`
	Summary    string     `gorm:"type:text" json:"summary"`
	Quantity   int        `gorm:"not null" json:"quantity"`
	CoverImage string     `gorm:"type:varchar(255)" json:"coverImage"`          // 封面圖片 URL
	Infos      []BookInfo `gorm:"foreignKey:BookId;references:ID" json:"infos"` // 一對多關聯
}

type BookInfo struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	BookId          string    `gorm:"not null" json:"book_id"`
	SerialNumber    string    `gorm:"type:varchar(100);not null" json:"serialNumber"`
	LastLendingDate time.Time `gorm:"column:lending_date" json:"lastLendingDate"`
	InStock         bool      `gorm:"default:true" json:"inStock"`

	Book Book `gorm:"foreignKey:BookId;references:ID" json:"book"`
}

type Category struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null;unique" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type BookCategory struct {
	BookID     string `gorm:"primaryKey" json:"book_id"`
	CategoryID string `gorm:"primaryKey" json:"category_id"`
}

type Review struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	BookID    string    `gorm:"not null" json:"book_id"`
	UserID    string    `gorm:"not null" json:"user_id"`
	Rating    int       `gorm:"not null" json:"rating"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `json:"created_at"`

	Book Book       `gorm:"foreignKey:BookID;references:ID"`
	User users.User `gorm:"foreignKey:UserID;references:ID"`
}
