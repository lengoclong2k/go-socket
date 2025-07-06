package usecases

import (
	"errors"
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"go-chat-app/internal/domain/repositories"
	"go-chat-app/internal/interfaces/dto"
	"time"
)

type RoomUseCase struct {
	roomRepo repositories.RoomRepository
	userRepo repositories.UserRepository
}

func NewRoomUseCase(roomRepo repositories.RoomRepository, userRepo repositories.UserRepository) *RoomUseCase {
	return &RoomUseCase{
		roomRepo: roomRepo,
		userRepo: userRepo,
	}
}

func (uc *RoomUseCase) CreateRoom(req dto.CreateRoomRequest, createdBy uuid.UUID) (*entities.Room, error) {
	room := &entities.Room{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		IsPrivate:   req.IsPrivate,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := uc.roomRepo.Create(room); err != nil {
		return nil, err
	}

	// Add creator as member
	if err := uc.roomRepo.AddMember(room.ID, createdBy); err != nil {
		return nil, err
	}

	return room, nil
}

func (uc *RoomUseCase) GetUserRooms(userID uuid.UUID) ([]entities.Room, error) {
	return uc.roomRepo.GetByUserID(userID)
}

func (uc *RoomUseCase) JoinRoom(roomID, userID uuid.UUID) error {
	room, err := uc.roomRepo.GetByID(roomID)
	if err != nil {
		return errors.New("room not found")
	}

	if room.IsPrivate {
		return errors.New("cannot join private room")
	}

	isMember, err := uc.roomRepo.IsUserMember(roomID, userID)
	if err != nil {
		return err
	}

	if isMember {
		return errors.New("user already member of room")
	}

	return uc.roomRepo.AddMember(roomID, userID)
}

func (uc *RoomUseCase) LeaveRoom(roomID, userID uuid.UUID) error {
	return uc.roomRepo.RemoveMember(roomID, userID)
}
