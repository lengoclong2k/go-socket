package repositories

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
)

type RoomRepository interface {
	Create(room *entities.Room) error
	GetByID(id uuid.UUID) (*entities.Room, error)
	GetByUserID(userID uuid.UUID) ([]entities.Room, error)
	Update(room *entities.Room) error
	Delete(id uuid.UUID) error
	AddMember(roomID, userID uuid.UUID) error
	RemoveMember(roomID, userID uuid.UUID) error
	GetMembers(roomID uuid.UUID) ([]entities.User, error)
	IsUserMember(roomID, userID uuid.UUID) (bool, error)
}
