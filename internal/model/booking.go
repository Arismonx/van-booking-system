package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	BookingStatusPending   = "pending"
	BookingStatusConfirmed = "confirmed"
	BookingStatusCancelled = "cancelled"
	BookingStatusCompleted = "completed"
)

const (
	PaymentStatusPending  = "pending"
	PaymentStatusPaid     = "paid"
	PaymentStatusFailed   = "failed"
	PaymentStatusRefunded = "refunded"
)

type (
	Booking struct {
		ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id"`
		NumSeats       uint32    `gorm:"not null" json:"num_seats"`
		TotalPrice     float64   `gorm:"type:numeric(10,2);not null" json:"total_price"`
		BookingStatus  string    `gorm:"type:varchar(50);not null;default:'pending'" json:"booking_status"`
		PaymentStatus  string    `gorm:"type:varchar(50);not null;default:'pending'" json:"payment_status"`
		BookingRefCode string    `gorm:"type:varchar(20);not null;uniqueIndex" json:"booking_ref_code"`
		CreatedAt      time.Time `gorm:"not null;autoCreateTime;" json:"created_at"`
		UpdatedAt      time.Time `gorm:"not null;autoUpdateTime;" json:"updated_at"`

		ScheduleID uuid.UUID `gorm:"type:uuid;not null;" json:"schedule_id"`
		Schedule   Schedule  `gorm:"foreignKey;ScheduleID;references:ID" json:"-"`

		UserID uuid.UUID `gorm:"type:uuid;not null;" json:"user_id"`
		User   User      `gorm:"foreignKey;UserID;references:ID" json:"-"`

		Payment *Payment `gorm:"foreignKey:BookingID" json:"payment,omitempty"`
	}
)
