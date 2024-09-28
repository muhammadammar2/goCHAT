package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	RoomType    string `json:"room_type"`           
	RoomCode    string `json:"room_code,omitempty"` 
	OwnerID     uint   `json:"owner_id"` 
    Owner       User   `gorm:"foreignKey:OwnerID"` 
}