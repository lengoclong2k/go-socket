package dto

import "github.com/google/uuid"

type SendMessageRequest struct {
	Content string    `json:"content" binding:"required"`
	RoomID  uuid.UUID `json:"room_id" binding:"required"`
}

type GetMessagesRequest struct {
	RoomID uuid.UUID `form:"room_id" binding:"required"`
	Limit  int       `form:"limit"`
	Offset int       `form:"offset"`
}
