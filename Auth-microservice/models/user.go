package models
import (
	"time"
)

// User represents the user model
type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    
	Email        string    `gorm:"not null;uniqueIndex"`
	Phone        string    `gorm:"not null;uniqueIndex"`
	DateOfBirth  string    
	PasswordHash string    `gorm:"not null"`
	RoleID       uint      `gorm:"not null"`
	Role         Role      `gorm:"foreignKey:RoleID"`
	AccessToken  string    
	RefreshToken string    
	ExpiresIn    time.Time 
	ExpiresAt    time.Time 
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}