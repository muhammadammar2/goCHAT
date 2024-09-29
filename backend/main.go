package main

import (
	"log"
	"os"

	"github.com/muhammadammar2/goCHAT/config"
	"github.com/muhammadammar2/goCHAT/middlewares"
	"github.com/muhammadammar2/goCHAT/redis"
	"github.com/muhammadammar2/goCHAT/utils"
	websockets "github.com/muhammadammar2/goCHAT/webSockets"

	"github.com/muhammadammar2/goCHAT/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {

	log.SetOutput(os.Stdout) 
	log.Println("Server is starting ...")	

	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middlewares.CORSMiddleware())

	db := config.ConnectDB()

	e.Validator = utils.NewValidator()

	redis.ConnectRedis()

	routes.SetupRoutes(e, db)

	go websockets.WebSocketHub.Run()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	
}