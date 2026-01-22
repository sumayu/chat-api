package repo

import (
 "context"

 "gorm.io/gorm"
 "yourmodule/internal/models"
)

type MessageRepo struct{ db *gorm.DB }

func NewMessageRepo(db *gorm.DB) *MessageRepo { return &MessageRepo{db: db} }

func (r *MessageRepo) Create(ctx context.Context, chatID int64, text string) (*models.Message, error) {
 msg := &models.Message{ChatID: chatID, Text: text}
 if err := r.db.WithContext(ctx).Create(msg).Error; err != nil {
  return nil, err
 }
 return msg, nil
}

func (r *MessageRepo) LastN(ctx context.Context, chatID int64, limit int) ([]models.Message, error) {
 var msgs []models.Message

 sub := r.db.WithContext(ctx).
  Model(&models.Message{}).
  Where("chat_id = ?", chatID).
  Order("created_at DESC").
  Limit(limit)

 if err := r.db.WithContext(ctx).
  Table("(?) as m", sub).
  Order("created_at ASC").
  Find(&msgs).Error; err != nil {
  return nil, err
 }
 return msgs, nil
}
