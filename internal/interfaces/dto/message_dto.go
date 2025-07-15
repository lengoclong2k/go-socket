package dto

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"time"
)

type SendMessageRequest struct {
	Content string    `json:"content" binding:"required"`
	RoomID  uuid.UUID `json:"room_id" binding:"required"`
}

type GetMessagesRequest struct {
	RoomID uuid.UUID `form:"room_id" binding:"required"`
	Limit  int       `form:"limit"`
	Offset int       `form:"offset"`
}

type MessageResponse struct {
	ID        uuid.UUID          `json:"id"`
	Content   string             `json:"content"`
	User      PublicUserResponse `json:"user"`
	RoomID    uuid.UUID          `json:"room_id"`
	CreatedAt time.Time          `json:"created_at"`
}

func ToMessageResponse(message *entities.Message, user *entities.User) MessageResponse {
	return MessageResponse{
		ID:        message.ID,
		Content:   message.Content,
		User:      ToPublicUserResponse(user),
		RoomID:    message.RoomID,
		CreatedAt: message.CreatedAt,
	}
}

// ToMessageResponses converts slice of entities.Message to slice of MessageResponse
func ToMessageResponses(messages []entities.Message, users map[uuid.UUID]*entities.User) []MessageResponse {
	responses := make([]MessageResponse, len(messages))
	for i, message := range messages {
		user := users[message.UserID]
		responses[i] = ToMessageResponse(&message, user)
	}
	return responses
}
