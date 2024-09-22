package handlers

import (
	"dummy/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

func CreateRoom(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        ownerID := uint(claims["userID"].(float64))

        var room struct {
            Name        string `json:"name" validate:"required"`
            Description string `json:"description"`
            Type        string `json:"type" validate:"required,oneof=public private"`
            Code        string `json:"code,omitempty"`
        }

        if err := c.Bind(&room); err != nil {
            return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
        }

        // new room
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

        return c.JSON(http.StatusOK, echo.Map{
            "room": newRoom,
        })
    }
}
