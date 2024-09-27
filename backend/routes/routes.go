package routes

import (
	"github.com/labstack/echo"
	"github.com/muhammadammar2/goCHAT/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	e.POST("/signup", func(c echo.Context) error {
		return handlers.Signup(c, db)
	})

	e.POST("/login" , func(c echo.Context) error {
		return handlers.Login(c , db)
	})
}
