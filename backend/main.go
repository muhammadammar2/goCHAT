package main

import (
	"log"
	"os"

	"github.com/muhammadammar2/goCHAT/config"

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

	db := config.ConnectDB()

	routes.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	
}