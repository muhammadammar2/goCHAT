package middlewares

import (
	redisclient "dummy/redis"
	"log"
	"net/http"
	"strings"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)
func JWTMiddleware(redisClient *redis.Client) echo.MiddlewareFunc {
    JWT_SECRET := "12hg3v1h23vh12v3h1v3gh12"
    if JWT_SECRET == "" {
        log.Fatal("JWT secret is missing")
    }

    return echojwt.WithConfig(echojwt.Config{
        SigningKey:  []byte(JWT_SECRET),
        TokenLookup: "header:Authorization",
        ContextKey:  "user",
        ErrorHandler: func(c echo.Context, err error) error {
            // Extract the Authorization header and remove "Bearer " prefix
            authHeader := c.Request().Header.Get("Authorization")
            token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

            if token == "" {
                log.Println("Authorization token is missing")
                return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization token")
            }

            // Check if the token is blacklisted
            blacklisted, redisErr := redisclient.IsTokenBlacklisted(redisClient, token)
            if redisErr != nil {
                log.Printf("Redis error while checking token blacklist: %v", redisErr)
                return echo.NewHTTPError(http.StatusInternalServerError, "Redis Error")
            }
            if blacklisted {
                log.Println("Token is blacklisted")
                return echo.ErrUnauthorized
            }

            // Log any JWT validation error
            log.Printf("JWT Error: %v", err)
            return echo.ErrUnauthorized
        },
    })
}