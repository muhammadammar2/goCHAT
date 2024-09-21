package models

import "time"

type Room struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Type        string `gorm:"size:10;not null"` // Can be "public" or "private"
	Code        string `gorm:"size:255"`         // Nullable, only for private rooms
	OwnerID     uint   `gorm:"not null"`         
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
