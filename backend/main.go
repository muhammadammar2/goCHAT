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

//     e.HTTPErrorHandler = func(err error, c echo.Context) {
//     log.Printf("Error: %v", err)
//     e.DefaultHTTPErrorHandler(err, c)
//   }


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
	jwtMiddleware := middlewares.JWTMiddleware()
	blacklistMiddleware := middlewares.TokenBlacklistMiddleware(redisClient)

	e.POST("/signup", handlers.Signup(db))
	e.POST("/login", handlers.Login(db))
	e.POST("/logout" , handlers.Logout(redisClient))
	
	r := e.Group("")
	r.Use(jwtMiddleware , blacklistMiddleware)
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		authHeader := c.Request().Header.Get("Authorization")
	// 		token := strings.TrimPrefix(authHeader, "Bearer ")

	// 		// Check if the token is blacklisted
	// 		blacklisted, err := redisclient.IsTokenBlacklisted(redisClient, token)
	// 		if err != nil {
	// 			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	// 		}
	// 		if blacklisted {
	// 			return echo.ErrUnauthorized
	// 		}

	// 		// Continue with the JWT middleware
	// 		return jwtMiddleware(next)(c)
	// 	}
	// })
	r.PUT("/update-profile" , handlers.UpdateProfile(db))
	r.DELETE("/delete", handlers.DeleteAccount(db))

	e.Logger.Fatal(e.Start(":8080"))
}
