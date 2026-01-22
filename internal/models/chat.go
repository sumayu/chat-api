package models

import "time"

type Chat struct {
 ID        int64     gorm:"primaryKey" json:"id"
 Title     string    gorm:"type:varchar(200);not null" json:"title"
 CreatedAt time.Time json:"created_at"
 Messages  []Message json:"-"
}
