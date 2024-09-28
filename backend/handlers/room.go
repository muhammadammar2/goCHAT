package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/muhammadammar2/goCHAT/models"
	"gorm.io/gorm"
)

func CreateRoom(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var req models.CreateRoomRequest

        if err := c.Bind(&req); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
        }

        if err := c.Validate(&req); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        }

        userID, ok := c.Get("user_id").(uint) 
        if !ok {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or missing user ID"})
        }

        room := models.Room{
            Name:        req.Name,
            Description: req.Description,
            RoomType:    req.RoomType,
            OwnerID:     userID,  
        }

		if req.RoomType == "private" && req.RoomCode == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Private rooms require a code"})
		}
		room.RoomCode = req.RoomCode

        if err := db.Create(&room).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create room"})
        }

        return c.JSON(http.StatusOK, room)
    }
}
