package database

import (
	"go-chat-app/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&entities.User{},
		&entities.Room{},
		&entities.Message{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
