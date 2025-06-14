package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusVanActive   = "active"
	StatusVanInactive = "inactive"
)

type (
	Van struct {
		ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id"`
		LicensePlate string    `gorm:"type:varchar(20);not null;uniqueIndex;" json:"license_plate"`
		Model        string    `gorm:"type:varchar(100);not null" json:"model"`
		Capacity     uint      `gorm:"not null" json:"capacity"`
		Status       string    `gorm:"type:varchar(50);not null;default:'active'" json:"status"`
		CreatedAt    time.Time `gorm:"not null;autoCreateTime;" json:"created_at"`
		UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;" json:"updated_at"`

		EmployeeID uuid.UUID `gorm:"type:uuid;not null" json:"employee_id"`
		Employee   Employee  `gorm:"foreignKey:EmployeeID;references:ID;" json:"-"`

		Schedules []Schedule `gorm:"foreignKey:VanID;" json:"schedules,omitempty"`
	}
)
