package handlers

import (
	"dummy/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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
		return echo.ErrUnauthorized
	   }
	   if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		  return echo.ErrUnauthorized
	   }

	   token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"email" : user.Email,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	   })

	   tokenString , err := token.SignedString([] byte (os.Getenv("JWT_SECRET")))
	   if err != nil {
		return err
	   }

	   return c.JSON(http.StatusOK , echo.Map {
		"token" : tokenString,
	   })
	}

}
 // not working bcs of the jwt
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

        if err := db.Where("email = ?", email).Delete(&models.User{}).Error; err != nil {
            return echo.ErrInternalServerError
        }

        return c.JSON(http.StatusOK, echo.Map{
            "message": "Account deleted successfully",
        })
    }
}


func Logout() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, echo.Map{
            "message": "Logout successful. clear your token on the client side",
        })
    }
}

