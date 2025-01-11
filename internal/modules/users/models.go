package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string         `gorm:"primaryKey"`
	AccountName string         `gorm:"type:varchar(50);unique;not null"`
	FirstName   string         `gorm:"type:varchar(150);not null"`
	LastName    string         `gorm:"type:varchar(150);not null"`
	Email       string         `gorm:"unique;not null"`
	Password    string         `gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Login struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type Role struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null;unique" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserRole struct {
	UserID string `gorm:"primaryKey" json:"user_id"`
	RoleID string `gorm:"primaryKey" json:"role_id"`
}

type SearchUserRequest struct {
	Email       string `json:"email"`
	AccountName string `json:"account_name"`
}
