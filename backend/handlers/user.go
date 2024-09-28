package handlers

// github.com/dgrijalva/jwt-go
import (
	"context"
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

// func GetUserProfile(db *gorm.DB) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         userID := c.Get("user_id")
//         if userID == nil {
//             return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User not authenticated"})
//         }

//         user, err := repository.GetUserByID(db, userID.(string)) 
//         if err != nil {
//             return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch profile"})
//         }

//         return c.JSON(http.StatusOK, user)
//     }
// }


// func UpdateProfile(db *gorm.DB) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         user := c.Get("user").(*jwt.Token)
//         claims, ok := user.Claims.(jwt.MapClaims)
//         if !ok || !user.Valid {
//             return echo.ErrUnauthorized
//         }

//         email, ok := claims["email"].(string)
//         if !ok {
//             return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
//         }

//         var existingUser models.User
//         if err := db.Where("email = ?", email).First(&existingUser).Error; err != nil {
//             return echo.NewHTTPError(http.StatusNotFound, "User not found")
//         }

//         updatedData := new(models.User)
//         if err := c.Bind(updatedData); err != nil {
//             return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
//         }

//         if updatedData.Username != "" && updatedData.Username != existingUser.Username {
//             if err := db.Where("username = ?", updatedData.Username).First(&models.User{}).Error; err == nil {
//                 return echo.NewHTTPError(http.StatusConflict, "Username already taken")
//             }
//             existingUser.Username = updatedData.Username
//         }

//         existingUser.Name = updatedData.Name

//         if updatedData.Password != "" {
//             hashPass, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
//             if err != nil {
//                 return echo.NewHTTPError(http.StatusInternalServerError, "Error hashing password")
//             }
//             existingUser.Password = string(hashPass)
//         }

//         if err := db.Save(&existingUser).Error; err != nil {
//             return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update profile")
//         }

//         return c.JSON(http.StatusOK, existingUser)
//     }
// }
