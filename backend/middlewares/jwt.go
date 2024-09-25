package middlewares

import (
	redisclient "dummy/redis"
	"log"
	"net/http"
	"os"
	"strings"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func JWTMiddleware(redisClient *redis.Client) echo.MiddlewareFunc {
	JWT_SECRET := os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatal("JWT secret is missing")
	}

	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(JWT_SECRET),
		TokenLookup: "header:Authorization",
		ContextKey:  "user",
		ErrorHandler: func(c echo.Context, err error) error {
			authHeader := c.Request().Header.Get("Authorization")
			log.Printf("Authorization header: %s", authHeader)

			token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			log.Printf("Processed token: %s", token)

			if token == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization token")
			}

			blacklisted, redisErr := redisclient.IsTokenBlacklisted(redisClient, token)
			if redisErr != nil {
				log.Printf("Redis Error: %v", redisErr)
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}
			if blacklisted {
				log.Printf("Token is blacklisted: %s", token)
				return echo.ErrUnauthorized
			}

			log.Printf("JWT Error: %v", err)
			return echo.ErrUnauthorized
		},
	})
}
