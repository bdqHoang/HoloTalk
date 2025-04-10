package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    
	Email        string    `gorm:"not null;unique_index"`
	Phone        string    `gorm:"not null;unique_index"`
	DateOfBirth  string
	PasswordHash string    `gorm:"not null"`
	RoleID       uint      `gorm:"not null"`
	Role         Role      `gorm:"foreignKey:RoleID"`
	AccessToken  string
	RefreshToken string
	ExpiresIn    time.Time
	ExpiresAt    time.Time
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}