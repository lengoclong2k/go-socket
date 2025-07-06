package entities

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	RoomID    uuid.UUID `json:"room_id" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User `json:"user" gorm:"foreignKey:UserID"`
	Room Room `json:"room" gorm:"foreignKey:RoomID"`
}
