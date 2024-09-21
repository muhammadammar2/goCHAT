package handlers

import (
	"dummy/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

// func CreateRoom(db *gorm.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(jwt.MapClaims)

// 		ownerID, ok := claims["userID"].(float64)
// 		log.Println("Owner ID:", ownerID)

// 		if !ok {
// 			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
// 		}

// 		var room struct {
// 			Name        string `json:"name" validate:"required"`
// 			Description string `json:"description"`
// 			Type        string `json:"type" validate:"required,oneof=public private"`
// 			Code        string `json:"code,omitempty"`
// 		}
// 		if err := c.Bind(&room); err != nil {
// 			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input", "details": err.Error()})
// 		}

// 		//new room
// 		newRoom := models.Room{
// 			Name:        room.Name,
// 			Description: room.Description,
// 			Type:        room.Type,
// 			OwnerID:     uint(ownerID),
// 			CreatedAt:   time.Now(),
// 			UpdatedAt:   time.Now(),
// 		}

// 		if err := db.Create(&newRoom).Error; err != nil {
// 			log.Println("Error creating room:", err)
// 			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Couldn't create room", "details": err.Error()})
// 		}
// 		return c.JSON(http.StatusOK, newRoom)
// 	}
// }


func CreateRoom(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        ownerID := uint(claims["userID"].(float64)) // Ensure you're getting the userID correctly

        var room struct {
            Name        string `json:"name" validate:"required"`
            Description string `json:"description"`
            Type        string `json:"type" validate:"required,oneof=public private"`
            Code        string `json:"code,omitempty"`
        }
        
        if err := c.Bind(&room); err != nil {
            return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
        }

        // New room
        newRoom := models.Room{
            Name:        room.Name,
            Description: room.Description,
            Type:        room.Type,
            OwnerID:     ownerID,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }

        if err := db.Create(&newRoom).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Couldn't create room"})
        }

        // Optionally generate a new token if needed (usually not necessary)
        tokenString, err := GenerateToken(ownerID, claims["email"].(string))
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
        }

        return c.JSON(http.StatusOK, echo.Map{
            "room": newRoom,
            "token": tokenString, // Include the new token if you want
        })
    }
}

