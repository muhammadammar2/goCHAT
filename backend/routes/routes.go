package routes

import (
	"github.com/labstack/echo"
	"github.com/muhammadammar2/goCHAT/handlers"
	"github.com/muhammadammar2/goCHAT/middlewares"
	"github.com/muhammadammar2/goCHAT/redis"

	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	e.POST("/signup", func(c echo.Context) error {
		return handlers.Signup(c, db)
	})

	e.POST("/login" , func(c echo.Context) error {
		return handlers.Login(c , db)
	})

	r := e.Group("")

	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.BlacklistMiddleware(redis.RedisClient))


	

	// r.PUT("/update-profile", handlers.UpdateProfile(db))
    r.POST("/create-room", handlers.CreateRoom(db))
    // r.GET("/profile", handlers.GetUserProfile(db)) 
    // r.DELETE("/delete", handlers.DeleteAccount(db))
	r.POST("/logout" , handlers.Logout)
	r.GET("/rooms", handlers.GetRooms(db))         
    r.POST("/join-room", handlers.JoinRoom(db)) 

	e.GET("/ws", handlers.WebSocketHandler)
}
