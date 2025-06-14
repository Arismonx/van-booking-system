package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	RoleUser     = "user"
	RoleEmployee = "employee"
	RoleAdmin    = "admin"
)

type (
	User struct {
		ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id"`
		Email        string    `gorm:"type:varchar(255);uniqueIndex;not null;" json:"email"`
		PasswordHash string    `gorm:"type:text;not null;" json:"-"`
		FirstName    string    `gorm:"type:varchar(255);not null;" json:"first_name"`
		LastName     string    `gorm:"type:varchar(255);not null;" json:"last_name"`
		PhoneNumber  *string   `gorm:"type:varchar(20);uniqueIndex;not null;" json:"phone_number"`
		Role         string    `gorm:"type:varchar(50);not null;default:'user'" json:"role"`
		CreatedAt    time.Time `gorm:"not null;autoCreateTime;" json:"created_at"`
		UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;" json:"updated_at"`

		Bookings []Booking `gorm:"foreignKey:UserID;" json:"booking,omitempty"`
	}
)
