package repository

import (
	"github.com/muhammadammar2/goCHAT/models"
	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, id string) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}