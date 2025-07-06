package entities

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	IsPrivate   bool      `json:"is_private" gorm:"default:false"`
	CreatedBy   uuid.UUID `json:"created_by" gorm:"type:uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Creator  User      `json:"creator" gorm:"foreignKey:CreatedBy"`
	Members  []User    `json:"members" gorm:"many2many:room_members;"`
	Messages []Message `json:"messages"`
}
