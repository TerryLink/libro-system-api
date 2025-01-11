package audit

import "time"

type AuditLog struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    *string   `json:"user_id"` // 可選，記錄執行操作的用戶
	Action    string    `gorm:"type:varchar(100);not null" json:"action"`
	Details   string    `gorm:"type:text" json:"details"` // 具體操作描述
	CreatedAt time.Time `json:"created_at"`
}
