package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusSchedule = "scheduled"
)

type (
	Schedule struct {
		ID                   uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id"`
		DepartureDate        time.Time  `gorm:"type:date;not null;" json:"departure_date"`
		DepartureTime        time.Time  `gorm:"type:time;not null;" json:"departure_time"`
		EstimatedArrivalTime *time.Time `gorm:"type:time;not null;" json:"estimated_arrival_time"`
		PricePerSeat         float64    `gorm:"type:numeric(10,2);not null;" json:"price_per_seat"`
		AvailableSeats       uint       `gorm:"not null;" json:"available_seats"`
		TotalSeats           uint       `gorm:"not null;" json:"total_seats"`
		Status               string     `gorm:"type:varchar(50);default:'scheduled';not null;" json:"status"`
		CreatedAt            time.Time  `gorm:"not null;autoCreateTime;" json:"created_at"`
		UpdatedAt            time.Time  `gorm:"not null;autoUpdateTime;" json:"updated_at"`

		VanID uuid.UUID `gorm:"type:uuid;not null;" json:"van_id"`
		Vans  Van       `gorm:"foreignKey:VanID;references:ID;" json:"-"`

		RoutesID uuid.UUID `gorm:"type:uuid;not null;" json:"routes_id"`
		Routes   Routes    `gorm:"foreignKey:RoutesID;references:ID;" json:"-"`

		Bookings []Booking `gorm:"foreignKey:ScheduleID;" json:"booking,omitempty"`
	}
)
