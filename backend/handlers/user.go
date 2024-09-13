package handlers

import (
	"dummy/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
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


	   return c.JSON(http.StatusOK , echo.Map {
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


func Logout() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, echo.Map{
            "message": "Logout successful. clear your token on the client side",
        })
    }
}

