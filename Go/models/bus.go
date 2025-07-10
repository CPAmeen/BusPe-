package models

import (
	"time"

	"github.com/google/uuid"
)

type Bus struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BusNumber string    `gorm:"uniqueIndex;not null"`
	OwnerID   uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
}
