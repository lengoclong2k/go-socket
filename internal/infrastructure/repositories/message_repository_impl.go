package repositories

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"go-chat-app/internal/domain/repositories"
	"gorm.io/gorm"
)

type messageRepositoryImpl struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) repositories.MessageRepository {
	return &messageRepositoryImpl{db: db}
}

func (r *messageRepositoryImpl) Create(message *entities.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepositoryImpl) GetByRoomID(roomID uuid.UUID, limit, offset int) ([]entities.Message, error) {
	var messages []entities.Message
	err := r.db.Preload("User").Where("room_id = ?", roomID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&messages).Error
	return messages, err
}

func (r *messageRepositoryImpl) GetByID(id uuid.UUID) (*entities.Message, error) {
	var message entities.Message
	err := r.db.Preload("User").First(&message, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *messageRepositoryImpl) Update(message *entities.Message) error {
	return r.db.Save(message).Error
}

func (r *messageRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Message{}, "id = ?", id).Error
}

func (r *messageRepositoryImpl) GetRoomMessageCount(roomID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Message{}).Where("room_id = ?", roomID).Count(&count).Error
	return count, err
}
