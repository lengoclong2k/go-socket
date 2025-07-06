package repositories

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
)

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id uuid.UUID) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetByUsername(username string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uuid.UUID) error
	GetAll() ([]entities.User, error)
	UpdateOnlineStatus(userID uuid.UUID, isOnline bool) error
}
