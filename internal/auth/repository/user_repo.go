package repository

import (
	"chat/internal/auth/model"
	mygorm "chat/internal/auth/repository/gorm"
	"context"
	"time"
)

type UserRepository struct {

	// Create(ctx context.Context, user *model.User) error
	// FindByID(ctx context.Context, id uint) (*model.User, error)
	// FindByUsername(ctx context.Context, username string) (*model.User, error)
	// FindByEmail(ctx context.Context, email string) (*model.User, error)
	// UpdateLastLogin(ctx context.Context, id uint) error
	// Update(ctx context.Context, user *model.User) error
	// Delete(ctx context.Context, id uint) error
}

func (*UserRepository) CreateUser(ctx *context.Context, user *model.User) error {
	return mygorm.MysqlDB.Create(user).Error
}

func (*UserRepository) FindByUsername(ctx *context.Context, username string) (*model.User, error) {
	var user model.User
	err := mygorm.MysqlDB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*UserRepository) UpdateLastLogin(ctx *context.Context, username string) error {
	var user model.User
	err := mygorm.MysqlDB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return err
	}
	user.LastLoginAt = time.Now()
	return nil
}
