package middlewares

import (
	redisclient "dummy/redis"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

// // TokenBlacklistMiddleware checks if the token is blacklisted
// func TokenBlacklistMiddleware(redisClient *redis.Client) echo.MiddlewareFunc {
//     return func(next echo.HandlerFunc) echo.HandlerFunc {
//         return func(c echo.Context) error {
//             // Extract the Authorization header
//             authHeader := c.Request().Header.Get("Authorization")
//             if authHeader == "" {
//                 log.Println("Missing Authorization header")
//                 return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
//             }

//             // Trim the "Bearer " prefix from the token
//             token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
//             if token == "" {
//                 log.Println("Invalid or missing token")
//                 return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing token")
//             }

//             // Check if the token is blacklisted
//             blacklisted, err := redisclient.IsTokenBlacklisted(redisClient, token)
//             if err != nil {
//                 log.Printf("Redis error during blacklist check: %v", err)
//                 return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
//             }

//             if blacklisted {
//                 log.Printf("Blacklisted token used: %s", token)
//                 return echo.NewHTTPError(http.StatusUnauthorized, "Token is blacklisted")
//             }

//             // Log and proceed to the next middleware/handler
//             log.Printf("Token is valid and not blacklisted: %s", token)
//             return next(c)
//         }
//     }
// }

// TokenBlacklistMiddleware checks if the token is blacklisted
func TokenBlacklistMiddleware(redisClient *redis.Client) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // Extract the Authorization header and trim "Bearer " prefix
            authHeader := c.Request().Header.Get("Authorization")
            if authHeader == "" {
                log.Println("Missing Authorization header")
                return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
            }

            // Trim the "Bearer " prefix from the token
            token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
            if token == "" {
                log.Println("Invalid or missing token")
                return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing token")
            }

            // Check if the token is blacklisted
            blacklisted, err := redisclient.IsTokenBlacklisted(redisClient, token)
            if err != nil {
                log.Printf("Redis error during blacklist check: %v", err)
                return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
            }

            if blacklisted {
                log.Printf("Blacklisted token used: %s", token)
                return echo.NewHTTPError(http.StatusUnauthorized, "Token is blacklisted")
            }

            // Log and proceed to the next middleware/handler
            log.Printf("Token is valid and not blacklisted: %s", token)
            return next(c)
        }
    }
}