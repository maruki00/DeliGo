package models

import (
	"time"
)

type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string    `gorm:"type:varchar(36);not null;index:idx_user" json:"user_id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Body      string    `gorm:"type:text;not null" json:"body"`
	IsRead    bool      `gorm:"type:boolean;default:false" json:"is_read"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type ChatMessage struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID     string    `gorm:"type:varchar(36);not null;index:idx_order_chat" json:"order_id"`
	SenderID    string    `gorm:"type:varchar(36);not null" json:"sender_id"`
	RecipientID string    `gorm:"type:varchar(36);not null" json:"recipient_id"`
	Message     string    `gorm:"type:text;not null" json:"message"`
	SentAt      time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"sent_at"`
}

type File struct {
	ID        string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	OwnerID   string    `gorm:"type:varchar(36);not null" json:"owner_id"`
	FileURL   string    `gorm:"type:varchar(512);not null" json:"file_url"`
	FileType  string    `gorm:"type:varchar(100);not null" json:"file_type"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
}
