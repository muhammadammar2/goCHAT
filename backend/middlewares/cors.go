package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// func CORSMiddleware() echo.MiddlewareFunc {
//     return middleware.CORSWithConfig(middleware.CORSConfig{
//         AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
//         AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
//         AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
//         AllowCredentials: true,
//         MaxAge: 300,
//     })
// }
func CORSMiddleware() echo.MiddlewareFunc {
    return middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"}, // Allow all origins (only for testing)
        AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
        AllowCredentials: true,
        MaxAge: 300,
    })
}
