package handlers

// github.com/dgrijalva/jwt-go
import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/muhammadammar2/goCHAT/models"
	"github.com/muhammadammar2/goCHAT/redis"
	"github.com/muhammadammar2/goCHAT/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

func Signup (c echo.Context , db * gorm.DB) error  {
	var signupRequest models.SignupRequest

	if err := c.Bind(&signupRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := validate.Struct(&signupRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	hashedPassword, err := utils.HashPassword(signupRequest.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	user := models.User{
		Username: signupRequest.Username,
		Email:    signupRequest.Email,
		Name : signupRequest.Name,
		Password: hashedPassword,
	}

	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully!"})

} 


func Login(c echo.Context, db *gorm.DB) error {
    var loginRequest models.LoginRequest

    if err := c.Bind(&loginRequest); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    if err := validate.Struct(&loginRequest); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
    }

    var user models.User

    if err := db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
    }

    tokenString, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Login successful",
        "token":   tokenString,
    })
}


var ctx = context.Background()

func Logout(c echo.Context) error {
    tokenString := c.Request().Header.Get("Authorization")
    if tokenString == "" {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
    }

    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    fmt.Println("Token String for logout:", tokenString)

    claims, err := utils.VerifyJWT(tokenString)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
    }

    exp := int64(claims.ExpiresAt)
    expirationTime := time.Unix(exp, 0)
    tokenTTL := time.Until(expirationTime)

    err = redis.RedisClient.Set(ctx, tokenString, "blacklisted", tokenTTL).Err()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to blacklist token"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
