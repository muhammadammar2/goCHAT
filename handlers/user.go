package handlers

import (
	"dummy/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

func Signup (c echo.Context) error {
	user := new (models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string (hash)
    
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated , user)
}

