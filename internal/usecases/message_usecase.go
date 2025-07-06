package usecases

import (
	"errors"
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"go-chat-app/internal/domain/repositories"
	"go-chat-app/internal/interfaces/dto"
	"time"
)

type MessageUseCase struct {
	messageRepo repositories.MessageRepository
	roomRepo    repositories.RoomRepository
}

func NewMessageUseCase(messageRepo repositories.MessageRepository, roomRepo repositories.RoomRepository) *MessageUseCase {
	return &MessageUseCase{
		messageRepo: messageRepo,
		roomRepo:    roomRepo,
	}
}

func (uc *MessageUseCase) SendMessage(req dto.SendMessageRequest, userID uuid.UUID) (*entities.Message, error) {
	// Check if user is member of room
	isMember, err := uc.roomRepo.IsUserMember(req.RoomID, userID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, errors.New("user is not a member of this room")
	}

	message := &entities.Message{
		ID:        uuid.New(),
		Content:   req.Content,
		UserID:    userID,
		RoomID:    req.RoomID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.messageRepo.Create(message); err != nil {
		return nil, err
	}

	// Get message with user info
	return uc.messageRepo.GetByID(message.ID)
}

func (uc *MessageUseCase) GetRoomMessages(roomID uuid.UUID, userID uuid.UUID, limit, offset int) ([]entities.Message, error) {
	// Check if user is member of room
	isMember, err := uc.roomRepo.IsUserMember(roomID, userID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, errors.New("user is not a member of this room")
	}

	return uc.messageRepo.GetByRoomID(roomID, limit, offset)
}
