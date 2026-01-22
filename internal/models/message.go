package models

import "time"

type Message struct {
 ID        int64     gorm:"primaryKey" json:"id"
 ChatID    int64     gorm:"not null;index" json:"chat_id"
 Text      string    gorm:"type:varchar(5000);not null" json:"text"
 CreatedAt time.Time json:"created_at"
}
