package dto

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"time"
)

type CreateRoomRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`
}

type JoinRoomRequest struct {
	RoomID uuid.UUID `json:"room_id" binding:"required"`
}

type RoomResponse struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	IsPrivate   bool                 `json:"is_private"`
	CreatedAt   time.Time            `json:"created_at"`
	CreatedBy   PublicUserResponse   `json:"created_by"`
	Members     []PublicUserResponse `json:"members"`
}

func ToRoomResponse(room *entities.Room, creator *entities.User, members []entities.User) RoomResponse {
	return RoomResponse{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		IsPrivate:   room.IsPrivate,
		CreatedBy:   ToPublicUserResponse(creator),
		CreatedAt:   room.CreatedAt,
		Members:     ToPublicUserResponses(members),
	}
}
