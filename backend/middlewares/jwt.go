// package middlewares
package middlewares

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/labstack/echo/v4"
// )

// // JWTMiddleware checks for a valid JWT in the Authorization header.
// func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			authHeader := c.Request().Header.Get("Authorization")
// 			if !strings.HasPrefix(authHeader, "Bearer ") {
// 				return c.JSON(http.StatusUnauthorized, "Missing or malformed JWT")
// 			}

// 			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 			// Parse the token
// 			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 				// Validate the signing method
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 				}
// 				return []byte(secretKey), nil
// 			})

// 			if err != nil || !token.Valid {
// 				return c.JSON(http.StatusForbidden, "Invalid or expired JWT")
// 			}

// 			// Set user info in context (optional)
// 			if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 				c.Set("user", claims["id"])
// 			}

// 			return next(c)
// 		}
// 	}
// }