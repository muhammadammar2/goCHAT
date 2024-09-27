package handlers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/muhammadammar2/goCHAT/models"
	"github.com/muhammadammar2/goCHAT/utils"
	"gorm.io/gorm"
)

var validate = validator.New()

func Signup (c echo.Context , db * gorm.DB) error  {
	var signupRequest models.SignupRequest

	if err := c.Bind(&signupRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := validate.Struct(&signupRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	hashedPassword, err := utils.HashPassword(signupRequest.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	user := models.User{
		Username: signupRequest.Username,
		Email:    signupRequest.Email,
		Name : signupRequest.Name,
		Password: hashedPassword,
	}

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully!"})

} 
