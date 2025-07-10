package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BusID       string    `gorm:"not null"`
	FareAmount  float64   `gorm:"not null"`
	PlatformFee float64   `gorm:"not null"`
	PaymentMode string    `gorm:"not null;default:'digital'"` // Optional for now
	PaidAt      time.Time `gorm:"autoCreateTime"`
}
