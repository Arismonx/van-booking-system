package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusOnLeave  = "on_leave"
)

type (
	Employee struct {
		ID                  uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id"`
		DriverLicenseNumber string     `gorm:"type:varchar(50);uniqueIndex;not null;" json:"driver_license_number"`
		EmployeeCode        *string    `gorm:"varchar(50);uniqueIndex;" json:"employee_code"`
		HireDate            *time.Time `gorm:"type:date;" json:"hire_date"`
		Status              string     `gorm:"type:varchr(50);not null;default:'active';" json:"status"`
		DateOfBirth         *time.Time `gorm:"type:date;" json:"date_of_birth"`
		Address             *string    `gorm:"type:text;" json:"address"`
		CreatedAt           time.Time  `gorm:"not null;autoCreateTime;" json:"created_at"`
		UpdatedAt           time.Time  `gorm:"not null;autoUpdateTime;" json:"updated_at"`

		UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex;" json:"user_id"`
		User   User      `gorm:"foreignKey;UserID;references:ID" json:"-"`

		Vans []Van `gorm:"foreignKey:EmployeeID;" json:"van,omitempty"`
	}
)
