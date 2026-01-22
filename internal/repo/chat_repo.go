package repo

import (
 "context"

 "gorm.io/gorm"
 "yourmodule/internal/models"
)

type ChatRepo struct{ db *gorm.DB }

func NewChatRepo(db *gorm.DB) *ChatRepo { return &ChatRepo{db: db} }

func (r *ChatRepo) Create(ctx context.Context, title string) (*models.Chat, error) {
 chat := &models.Chat{Title: title}
 if err := r.db.WithContext(ctx).Create(chat).Error; err != nil {
  return nil, err
 }
 return chat, nil
}

func (r *ChatRepo) GetByID(ctx context.Context, id int64) (*models.Chat, error) {
 var chat models.Chat
 if err := r.db.WithContext(ctx).First(&chat, id).Error; err != nil {
  return nil, err
 }
 return &chat, nil
}

func (r *ChatRepo) Delete(ctx context.Context, id int64) error {
 return r.db.WithContext(ctx).Delete(&models.Chat{}, id).Error
}
