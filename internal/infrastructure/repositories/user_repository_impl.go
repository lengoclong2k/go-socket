package repositories

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"go-chat-app/internal/domain/repositories"
	"gorm.io/gorm"
	"time"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) GetByID(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *userRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.User{}, "id = ?", id).Error
}

func (r *userRepositoryImpl) GetAll() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepositoryImpl) UpdateOnlineStatus(userID uuid.UUID, isOnline bool) error {
	return r.db.Model(&entities.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"is_online": isOnline,
		"last_seen": time.Now(),
	}).Error
}
