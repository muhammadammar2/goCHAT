package middlewares

import (
	redisclient "dummy/redis"
	"log"
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
                    log.Println("Redis error during blacklist check:", err)
                    return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
                }
                if blacklisted {
                    log.Println("Blacklisted token used:", token)
                    return echo.ErrUnauthorized
                }
                

                return next(c)
            }
        }
    }