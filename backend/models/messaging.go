package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	RoomID  uint   `json:"room_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}
