package middlewares

import (
	"log"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	JWT_SECRET := "12hg3v1h23vh12v3h1v3gh12"
	if JWT_SECRET == "" {
		log.Fatal("prob in JWT")
	}

	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(JWT_SECRET),
		TokenLookup: "header:Authorization",
		ContextKey:  "user",
		ErrorHandler: func(c echo.Context, err error) error {
			log.Printf("JWT Error: %v", err)
			return echo.ErrUnauthorized
		},
	})
}
