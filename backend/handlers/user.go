// package handlers

// import (
// 	"dummy/models"
// 	redisclient "dummy/redis"
// 	"net/http"
// 	"os"

// 	"strings"
// 	"time"

// 	// "github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// 	echojwt "github.com/labstack/echo-jwt"

// 	"github.com/redis/go-redis/v9"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// func Signup(db *gorm.DB) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         // body, err := io.ReadAll(c.Request().Body)
//         // if err != nil {
//         //     log.Printf("Eror reading request body: %v", err)
//         //     return echo.NewHTTPError(http.StatusBadRequest, "Error reading request")
//         // }
//         // c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

//         signupData := struct {
//             Username string `json:"username"`
//             Name     string `json:"name"`
//             Email    string `json:"email"`
//             Password string `json:"password"`
//         }{}

//         if err := c.Bind(&signupData); err != nil {
//             return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
//         }

//         if signupData.Password == "" {
//             return echo.NewHTTPError(http.StatusBadRequest, "Password is required")
//         }

//         hashPass, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
//         if err != nil {
//             return echo.NewHTTPError(http.StatusInternalServerError, "Error processing password")
//         }

//         user := &models.User{
//             Username: signupData.Username,
//             Name:     signupData.Name,
//             Email:    signupData.Email,
//             Password: string(hashPass),
//         }

//         if err := db.Create(user).Error; err != nil {
//             return echo.NewHTTPError(http.StatusInternalServerError, "Error creating user")
//         }

//         return c.JSON(http.StatusCreated, echo.Map{
//             "message": "User created successfully",
//             "userId":  user.ID,
//         })
//     }
// }

// // func Login(db *gorm.DB) echo.HandlerFunc {
// //     return func(c echo.Context) error {
// //         login := new(struct {
// //             Email    string `json:"email"`
// //             Password string `json:"password"`
// //         })

// //         if err := c.Bind(login); err != nil {
// //             return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
// //         }
// //         user := new(models.User)
// //         if err := db.Where("email = ?", login.Email).First(user).Error; err != nil {
// //             if err == gorm.ErrRecordNotFound {
// //                 return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
// //             }
// //             return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
// //         }

// //         err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
// //         if err != nil {
// //             return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
// //         }

// //         JWT_SECRET := "3287y48327483264b32864cc3b462c3"
// //         token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// //             "userID": user.ID,
// //             "email":  user.Email,
// //             "exp":    time.Now().Add(time.Hour * 24).Unix(),
// //         })

// //         tokenString, err := token.SignedString([]byte(JWT_SECRET))
// //         if err != nil {
// //             return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
// //         }

// //         return c.JSON(http.StatusOK, echo.Map{
// //             "token": tokenString,
// //         })
// //     }
// // }

// func GenerateToken(userID uint, email string) (string, error) {
//     JWT_SECRET := os.Getenv("JWT_SECRET")
//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//         "userID": userID,
//         "email":  email,
//         "exp":    time.Now().Add(time.Hour * 24).Unix(),
//     })
//     return token.SignedString([]byte(JWT_SECRET))
// }

// func ValidateToken(c echo.Context) error {
// 	// Extract token from Authorization header
// 	tokenString := c.Request().Header.Get("Authorization")
// 	if tokenString == "" {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "Missing token",
// 		})
// 	}

// 	// Parse and validate the token
// 	token, err := jwt.Parse(tokenString[len("Bearer "):], func(token *jwt.Token) (interface{}, error) {
// 		// Check if the token's signing method matches
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid signing method")
// 		}
// 		// Return the secret key
// 		return []byte("JWT_SECRET"), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "Invalid or expired token",
// 		})
// 	}

// 	// You can also check if the token is blacklisted (if using Redis-based token blacklisting)
// 	// Assuming you have a `CheckBlacklistedToken` function
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "Invalid token",
// 		})
// 	}

// 	// Check if token is expired
// 	expiration := claims["exp"].(float64)
// 	if time.Unix(int64(expiration), 0).Before(time.Now()) {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "Token is expired",
// 		})
// 	}

// 	// Optional: Check if token is blacklisted
// 	// if CheckBlacklistedToken(tokenString) {
// 	//     return c.JSON(http.StatusUnauthorized, map[string]string{
// 	//         "message": "Token has been blacklisted",
// 	//     })
// 	// }

// 	// If the token is valid, return success response
// 	return c.JSON(http.StatusOK, map[string]string{
// 		"message": "Token is valid",
// 	})
// }

// func Login(db *gorm.DB) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         login := new(struct {
//             Email    string `json:"email"`
//             Password string `json:"password"`
//         })

//         if err := c.Bind(login); err != nil {
//             return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
//         }

//         user := new(models.User)
//         if err := db.Where("email = ?", login.Email).First(user).Error; err != nil {
//             if err == gorm.ErrRecordNotFound {
//                 return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
//             }
//             return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
//         }

//         err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
//         if err != nil {
//             return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
//         }

//         // Call GenerateToken to create a token for the user
//         tokenString, err := GenerateToken(user.ID, user.Email)
//         if err != nil {
//             return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
//         }

//         return c.JSON(http.StatusOK, echo.Map{
//             "token": tokenString,
//         })
//     }
// }

// func Logout(client *redis.Client) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         authHeader := c.Request().Header.Get("Authorization")
//         token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

//         if token == "" {
//             return echo.NewHTTPError(http.StatusBadRequest, "No token provided")
//         }
//         expiration := 24 * time.Hour

//         err := redisclient.BlacklistToken(client , token, expiration)
//         if err != nil {
//             return echo.NewHTTPError(http.StatusInternalServerError, "Could not blacklist token")
//         }
//         return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
//     }
// }

// func DeleteAccount(db *gorm.DB) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         user := c.Get("userID").(*jwt.Token)
//         claims, ok := user.Claims.(jwt.MapClaims)
//         if !ok {
//             return echo.ErrUnauthorized
//         }

//         email, ok := claims["email"].(string)
//         if !ok {
//             return echo.ErrUnauthorized
//         }

//         result := db.Unscoped().Where("email = ?", email).Delete(&models.User{})
//         if result.Error != nil {
//             return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
//         }
//         if result.RowsAffected == 0 {
//             return echo.NewHTTPError(http.StatusNotFound, "User not found")
//         }

//         return c.JSON(http.StatusOK, echo.Map{
//             "message": "Account deleted successfully",
//         })
//     }
// }

// func UpdateProfile(db *gorm.DB) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         user := c.Get("user").(*jwt.Token)
//         claims, ok := user.Claims.(jwt.MapClaims)
//         if !ok {
//             return echo.ErrUnauthorized
//         }

//         email, ok := claims["email"].(string)
//         if !ok {
//             return echo.ErrUnauthorized
//         }

//         var existingUser models.User
//         if err := db.Where("email = ?", email).First(&existingUser).Error; err != nil {
//             return echo.NewHTTPError(http.StatusNotFound, "User not found")
//         }

//         updatedData := new(models.User)
//         if err := c.Bind(updatedData); err != nil {
//             return err
//         }

//         if updatedData.Username != existingUser.Username {
//             var userWithSameUsername models.User
//             if err := db.Where("username = ?", updatedData.Username).First(&userWithSameUsername).Error; err == nil {
//                 return echo.NewHTTPError(http.StatusConflict, "Username already in use")
//             }
//         }

//         existingUser.Username = updatedData.Username
//         existingUser.Name = updatedData.Name

//         if updatedData.Password != "" {
//             hashPass, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
//             if err != nil {
//                 return err
//             }
//             existingUser.Password = string(hashPass)
//         }

//         if err := db.Save(&existingUser).Error; err != nil {
//             return err
//         }

//         return c.JSON(http.StatusOK, existingUser)
//     }
// }

// // // GetUserProfile handler to fetch user's profile
// // func GetUserProfile(db *gorm.DB) echo.HandlerFunc {
// //     return func(c echo.Context) error {
// //         log.Println("GetUserProfile handler called")

// //         // Retrieve the user (token) from the Echo context
// //         user := c.Get("user")
// //         if user == nil {
// //             log.Println("User not found in context")
// //             return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
// //         }

// //         // Parse the JWT token from context
// //         token, ok := user.(*jwt.Token)
// //         if !ok {
// //             log.Println("Failed to parse JWT token from context")
// //             return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
// //         }

// //         // Extract claims from the JWT token
// //         claims, ok := token.Claims.(jwt.MapClaims)
// //         if !ok {
// //             log.Println("Failed to parse claims from JWT token")
// //             return echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims")
// //         }

// //         // Extract username from claims
// //         username, ok := claims["username"].(string)
// //         if !ok || username == "" {
// //             log.Println("Username claim not found or invalid in JWT token")
// //             return echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims")
// //         }

// //         log.Printf("Fetching profile for user: %s", username)

// //         // Fetch the user profile from the database
// //         var profile models.User
// //         if err := db.Where("username = ?", username).First(&profile).Error; err != nil {
// //             log.Printf("Error fetching the user profile for username %s: %v", username, err)
// //             return echo.NewHTTPError(http.StatusNotFound, "User not found")
// //         }

// //         // Return the profile information as a JSON response
// //         return c.JSON(http.StatusOK, profile)
// //     }
// // }

package handlers

import (
	"dummy/models"
	redisclient "dummy/redis"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		signupData := struct {
			Username string `json:"username"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}

		if err := c.Bind(&signupData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
		}

		if signupData.Password == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Password is required")
		}

		hashPass, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error processing password")
		}

		user := &models.User{
			Username: signupData.Username,
			Name:     signupData.Name,
			Email:    signupData.Email,
			Password: string(hashPass),
		}

		if err := db.Create(user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error creating user")
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "User created successfully",
			"userId":  user.ID,
		})
	}
}

func GenerateToken(userID uint, email string) (string, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(JWT_SECRET))
}

func ValidateToken(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Missing token",
		})
	}

	token, err := jwt.Parse(tokenString[len("Bearer "):], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid or expired token",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Token is valid",
	})
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		login := new(struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		})

		if err := c.Bind(login); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
		}

		user := new(models.User)
		if err := db.Where("email = ?", login.Email).First(user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
		}

		tokenString, err := GenerateToken(user.ID, user.Email)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": tokenString,
		})
	}
}

func Logout(client *redis.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		if token == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "No token provided")
		}
		expiration := 24 * time.Hour

		err := redisclient.BlacklistToken(client, token, expiration)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Could not blacklist token")
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
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

