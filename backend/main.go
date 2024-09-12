package main

import (
	"dummy/handlers"
	"dummy/models"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)	

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

	log.Println("JWT_SECRET:", os.Getenv("JWT_SECRET")) // check the value of JWT_SECRET
	

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalf("JWT_SECRET not set in .env file")
	}

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(jwtSecret),
		TokenLookup:   "header:Authorization",
		ContextKey:    "user", 
		ErrorHandler: func(c echo.Context, err error) error {
			log.Printf("JWT Error: %v", err)
			return echo.ErrUnauthorized
		},
	})	
	//testing the jwt
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				c.Request().Header.Set("Authorization", strings.TrimPrefix(authHeader, "Bearer "))
			}
			return next(c)
		}
	})

	//custom middlware	
	// e.Use(middlewares.JWTMiddleware(jwtSecret))
	// e.GET("/public", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Public endpoint")
	// })
	// e.GET("/protected", func(c echo.Context) error {
	// 	user := c.Get("user") // Retrieve user info from context
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"user": user,
	// 	})
	// })

		

	e.POST("/signup", handlers.Signup(db))
	e.POST("/login", handlers.Login(db))
	e.POST("/logout" , handlers.Logout())
	e.GET("/test", func(c echo.Context) error {
		user := c.Get("user")
		if user == nil {
			return echo.ErrUnauthorized
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Token is valid"})
	})
	

	r := e.Group("")
	r.Use(jwtMiddleware)
	r.DELETE("/delete", handlers.DeleteAccount(db))

	e.Logger.Fatal(e.Start(":8080"))
}
