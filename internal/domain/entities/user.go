package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Avatar    string    `json:"avatar"`
	IsOnline  bool      `json:"is_online" gorm:"default:false"`
	LastSeen  time.Time `json:"last_seen"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
