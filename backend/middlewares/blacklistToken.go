package middlewares

import (
	redisclient "dummy/redis"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func TokenBlacklistMiddleware(redisClient *redis.Client) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            authHeader := c.Request().Header.Get("Authorization")
            token := strings.TrimPrefix(authHeader, "Bearer ")

            blacklisted, err := redisclient.IsTokenBlacklisted(redisClient, token)
            if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
            }
            if blacklisted {
                return echo.ErrUnauthorized
            }

            return next(c)
        }
    }
}