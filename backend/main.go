package main

import (
	"dummy/handlers"
	"dummy/models"
	"log"
	"net/http"
	"os"

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

	// log.Println("JWT_SECRET:", os.Getenv("JWT_SECRET")) // check the value of JWT_SECRET
	

	jwtSecret := os.Getenv("JWT_SECRET")
	JWT_SECRET := "12hg3v1h23vh12v3h1v3gh12"  

	if jwtSecret == "" {
		// log.Fatalf("JWT_SECRET not set in .env file" )
		log.Fatal("prob in JWT")
	}

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(JWT_SECRET),	
		TokenLookup:   "header:Authorization",
		ContextKey:    "user", 
		ErrorHandler: func(c echo.Context, err error) error {
			log.Printf("JWT Error: %v", err)
			return echo.ErrUnauthorized
		},
	})	
	
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
	
    // e.DELETE("/delete" , handlers.DeleteAccount(db))
	r := e.Group("")
	r.Use(jwtMiddleware)
	r.DELETE("/delete", handlers.DeleteAccount(db))

	e.Logger.Fatal(e.Start(":8080"))
}
