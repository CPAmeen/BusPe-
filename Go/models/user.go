package models

import "time"

type Users struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Phone     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}
