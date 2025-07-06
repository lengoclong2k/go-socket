package repositories

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
)

type MessageRepository interface {
	Create(message *entities.Message) error
	GetByRoomID(roomID uuid.UUID, limit, offset int) ([]entities.Message, error)
	GetByID(id uuid.UUID) (*entities.Message, error)
	Update(message *entities.Message) error
	Delete(id uuid.UUID) error
	GetRoomMessageCount(roomID uuid.UUID) (int64, error)
}
