package main

import (
	"dummy/handlers"
	"dummy/middlewares"
	"dummy/models"
	redisclient "dummy/redis"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)	

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.CORSMiddleware())

	log.SetOutput(os.Stdout)
	log.Println("Server starting...")

	if err := godotenv.Load(); err != nil {	
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
	
	 dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}	

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	redisClient := redisclient.NewRedisClient()
	jwtMiddleware := middlewares.JWTMiddleware(redisClient)
	blacklistMiddleware := middlewares.TokenBlacklistMiddleware(redisClient)

	e.POST("/signup", handlers.Signup(db))
	e.POST("/login", handlers.Login(db))
	e.POST("/logout" , handlers.Logout(redisClient))
	
	r := e.Group("")
	r.Use(jwtMiddleware , blacklistMiddleware)
	r.PUT("/update-profile" , handlers.UpdateProfile(db))
	r.POST("/create-room" , handlers.CreateRoom(db))
	// r.GET("/profile", handlers.GetUserProfile(db))
	r.DELETE("/delete", handlers.DeleteAccount(db))

	e.Logger.Fatal(e.Start(":8080"))
}
	