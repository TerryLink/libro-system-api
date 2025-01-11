package audit

import "time"

type AuditLog struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    *string   `json:"user_id"`
	Action    string    `gorm:"type:varchar(100);not null" json:"action"`
	Details   string    `gorm:"type:text" json:"details"`
	CreatedAt time.Time `json:"created_at"`
}
