package repositories

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"go-chat-app/internal/domain/repositories"
	"gorm.io/gorm"
)

type roomRepositoryImpl struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) repositories.RoomRepository {
	return &roomRepositoryImpl{db: db}
}

func (r *roomRepositoryImpl) Create(room *entities.Room) error {
	return r.db.Create(room).Error
}

func (r *roomRepositoryImpl) GetByID(id uuid.UUID) (*entities.Room, error) {
	var room entities.Room
	err := r.db.Preload("Creator").Preload("Members").First(&room, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepositoryImpl) GetByUserID(userID uuid.UUID) ([]entities.Room, error) {
	var rooms []entities.Room
	err := r.db.Preload("Creator").Preload("Members").
		Joins("JOIN room_members ON rooms.id = room_members.room_id").
		Where("room_members.user_id = ?", userID).
		Find(&rooms).Error
	return rooms, err
}

func (r *roomRepositoryImpl) Update(room *entities.Room) error {
	return r.db.Save(room).Error
}

func (r *roomRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Room{}, "id = ?", id).Error
}

func (r *roomRepositoryImpl) AddMember(roomID, userID uuid.UUID) error {
	return r.db.Exec("INSERT INTO room_members (room_id, user_id) VALUES (?, ?) ON CONFLICT DO NOTHING", roomID, userID).Error
}

func (r *roomRepositoryImpl) RemoveMember(roomID, userID uuid.UUID) error {
	return r.db.Exec("DELETE FROM room_members WHERE room_id = ? AND user_id = ?", roomID, userID).Error
}

func (r *roomRepositoryImpl) GetMembers(roomID uuid.UUID) ([]entities.User, error) {
	var users []entities.User
	err := r.db.Joins("JOIN room_members ON users.id = room_members.user_id").
		Where("room_members.room_id = ?", roomID).
		Find(&users).Error
	return users, err
}

func (r *roomRepositoryImpl) IsUserMember(roomID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Table("room_members").Where("room_id = ? AND user_id = ?", roomID, userID).Count(&count).Error
	return count > 0, err
}
