package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Payment struct {
		ID            uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id"`
		Amount        float64    `gorm:"type:numeric(10,2);not null;check:amount > 0" json:"amount"`
		PaymentMethod string     `gorm:"type:varchar(50);not null" json:"payment_method"`
		TransactionID *string    `gorm:"type:varchar(255);uniqueIndex" json:"transaction_id"`
		PaymentStatus string     `gorm:"type:varchar(50);default:'pending';not null" json:"payment_status"`
		PaidAt        *time.Time `gorm:"type:timestamptz" json:"paid_at"`

		BookingID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex;" json:"booking_id"`
		Booking   Booking   `gorm:"foreignKey;BookingID;references:ID" json:"-"`
	}
)
