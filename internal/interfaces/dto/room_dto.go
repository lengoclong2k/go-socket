package dto

import "github.com/google/uuid"

type CreateRoomRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`
}

type JoinRoomRequest struct {
	RoomID uuid.UUID `json:"room_id" binding:"required"`
}
