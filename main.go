package main

import (
	"dummy/handlers"
	"dummy/models"
	"log"

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

	if err:= godotenv.Load(); err != nil {
		log.Printf("no env found")
	}

	dsn := "root:123@tcp(127.0.0.1:3306)/goo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	e.POST("/signup" , handlers.Signup(db))
	e.POST("/login" , handlers.Login(db) )

	e.Logger.Fatal(e.Start(":8080"))
}
