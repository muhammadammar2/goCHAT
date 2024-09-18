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
            authHeader := c.Request().Header.Get("Authorization")
            token := strings.TrimPrefix(authHeader, "Bearer ")

            blacklisted, redisErr := redisclient.IsTokenBlacklisted(redisClient, token)
            if redisErr != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, "Redis Error")
            }
            if blacklisted {
                log.Println("Token is blacklisted")
                return echo.ErrUnauthorized
            }

            log.Printf("JWT Error: %v", err)
            return echo.ErrUnauthorized
        },
    })
}

