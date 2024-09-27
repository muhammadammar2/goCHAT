package middlewares

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func BlacklistMiddleware(redisClient *redis.Client) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            token := c.Request().Header.Get("Authorization")
            if token == "" {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
            }

            if isBlacklisted(redisClient, token) {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token has been blacklisted"})
            }

            return next(c)
        }
    }
}

func isBlacklisted(redisClient *redis.Client, token string) bool {
    val, err := redisClient.Get(ctx, token).Result()
    return err == nil && val == "blacklisted"
}