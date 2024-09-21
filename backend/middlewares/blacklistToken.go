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
            if authHeader == "" {
                return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
            }

            token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
            if token == "" {
                return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing token")
            }

            blacklisted, err := redisclient.IsTokenBlacklisted(redisClient, token)
            if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
            }

            if blacklisted {
                return echo.NewHTTPError(http.StatusUnauthorized, "Token is blacklisted")
            }

            return next(c)
        }
    }
}