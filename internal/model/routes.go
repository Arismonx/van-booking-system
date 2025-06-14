package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Routes struct {
		ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id"`
		Origin      string    `gorm:"type:varchar(100);not null;" json:"origin"`
		Destination string    `gorm:"type:varchar(100);not null;" json:"destination"`
		DistanceKM  *float64  `gorm:"type:numeric(10,2);" json:"distance_km"`
		CreatedAt   time.Time `gorm:"not null;autoCreateTime;" json:"created_at"`
		UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;" json:"updated_at"`

		Schedules []Schedule `gorm:"foreignKey:RoutesID;" json:"schedules,omitempty"`
	}
)
