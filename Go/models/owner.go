package models

import (
	"time"

	"github.com/google/uuid"
)

type Owner struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string
	UPIID     string `gorm:"not null"`
	CreatedAt time.Time
	Buses     []Bus `gorm:"foreignKey:OwnerID"`
}
