package model

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey;column:id;autoIncrement:true"`
	Username     string    `gorm:"column:username;type:varchar(50);uniqueIndex;not null"`
	Email        string    `gorm:"column:email;type:varchar(100);uniqueIndex;not null"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
	LastLoginAt  time.Time `gorm:"column:last_login_at"`
	IsActive     bool      `gorm:"column:is_active;type:tinyint(1);default:1"`
	AvatarURL    *string   `gorm:"column:avatar_url;type:varchar(255)"`
}
