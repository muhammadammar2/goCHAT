package middlewares

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

func BlacklistMiddleware(redisClient *redis.Client) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func (c echo.Context) error {
			token := c.Request().Header.Get("Authorization")

			exists, err := redisClient.Exists(context.Background(), token).Result()

			if err != nil {
                return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
            }

			if exists > 0 {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is blacklisted"})
            }

			return next(c)
		}
	}
}