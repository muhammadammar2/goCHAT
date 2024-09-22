package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//id aah
	Username string `gorm:"unique" json:"username"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}
