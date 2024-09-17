package handlers

import (
	"dummy/models"
	redisclient "dummy/redis"
	"log"
	"net/http"

	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func Signup (db * gorm.DB) echo.HandlerFunc{
	return func (c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return err
		}
	

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string (hashPass)
    
	if err := db.Create(user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated , user)
}

}

func Login (db * gorm.DB) echo.HandlerFunc {
	return func (c echo.Context) error  {   
       login := new (models.User)
	   if err := c.Bind(login); err != nil {
		return err
	   }

	   user := new (models.User)
	   if err := db.Where("email = ?" , login.Email).First(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized , "Invalid Email || pass")
	   }
	   if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		  return echo.NewHTTPError(http.StatusUnauthorized , "Invalid Email || pass")
	   }
	   JWT_SECRET := "12hg3v1h23vh12v3h1v3gh12"  

	   token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"userID" : user.ID,
		"email" : user.Email,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),	
	   })

	//    tokenString , err := token.SignedString([] byte (os.Getenv("JWT_SECRET")))
	tokenString , err := token.SignedString([] byte (JWT_SECRET)) 
	   if err != nil {
		return err
	   }


	   return c.JSON(http.StatusOK ,echo.Map {
		"token" : tokenString,
	   })
	}

}

func DeleteAccount(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims, ok := user.Claims.(jwt.MapClaims)
        if !ok {
            return echo.ErrUnauthorized
        }

        email, ok := claims["email"].(string)
        if !ok {
            return echo.ErrUnauthorized
        }

        result := db.Unscoped().Where("email = ?", email).Delete(&models.User{})
        if result.Error != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
        }
        if result.RowsAffected == 0 {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }

        return c.JSON(http.StatusOK, echo.Map{
            "message": "Account deleted successfully",
        })
    }
}

func Logout(client *redis.Client) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")
        token := strings.TrimPrefix(authHeader, "Bearer ")

        if token == "" {
            log.Println("No token provided for logout")
            return echo.NewHTTPError(http.StatusBadRequest, "No token provided")
        }

        expiration := 24 * time.Hour

        err := redisclient.BlacklistToken(client, token, expiration)
        if err != nil {
            log.Printf("Error during logout: %v", err)
            return echo.NewHTTPError(http.StatusInternalServerError, "Could not blacklist token")
        }

        log.Println("User logged out successfully")
        return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
    }
}

func UpdateProfile(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims, ok := user.Claims.(jwt.MapClaims)
        if !ok {
            return echo.ErrUnauthorized
        }

        email, ok := claims["email"].(string)
        if !ok {
            return echo.ErrUnauthorized
        }

        var existingUser models.User
        if err := db.Where("email = ?", email).First(&existingUser).Error; err != nil {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }

        updatedData := new(models.User)
        if err := c.Bind(updatedData); err != nil {
            return err
        }

        if updatedData.Username != existingUser.Username {
            var userWithSameUsername models.User
            if err := db.Where("username = ?", updatedData.Username).First(&userWithSameUsername).Error; err == nil {
                return echo.NewHTTPError(http.StatusConflict, "Username already in use")
            }
        }

        existingUser.Username = updatedData.Username
        existingUser.Name = updatedData.Name

        if updatedData.Password != "" {
            hashPass, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
            if err != nil {
                return err
            }
            existingUser.Password = string(hashPass)
        }

        if err := db.Save(&existingUser).Error; err != nil {
            return err
        }

        return c.JSON(http.StatusOK, existingUser)
    }
}


func GetUserProfile(db *gorm.DB, username string) (models.User, error) {
    var user models.User
    err := db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return user, err
    }
    return user, nil
}



func GetProfile(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims, ok := user.Claims.(jwt.MapClaims)
        if !ok {
            return echo.ErrUnauthorized
        }

        username, ok := claims["username"].(string)
        if !ok {
            return echo.ErrUnauthorized
        }

        profile, err := GetUserProfile(db, username)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Unable to retrieve user profile")
        }

        return c.JSON(http.StatusOK, profile)
    }
}
