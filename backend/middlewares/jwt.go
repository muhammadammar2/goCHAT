package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/muhammadammar2/goCHAT/utils"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
		}

		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := utils.VerifyJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		c.Set("user_id", claims.UserID)
		return next(c)
	}
}


// func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		token := c.Request().Header.Get("Authorization")
// 		if token == "" {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
// 		}

// 		token = strings.TrimPrefix(token, "Bearer ")
// 		if token == "" {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
// 		}

// 		claims, err := utils.VerifyJWT(token)
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
// 		}

// 		c.Set("user_id", claims.UserID)

// 		return next(c)
// 	}
// }