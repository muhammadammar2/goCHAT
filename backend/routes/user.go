package routes

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Define your routes here
	// e.POST("/signup", handlers.Signup(db))
	// e.POST("/login", handlers.Login(db))
}