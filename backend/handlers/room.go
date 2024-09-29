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


func GetRooms(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var rooms []models.Room
        if err := db.Find(&rooms).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve rooms"})
        }
        return c.JSON(http.StatusOK, rooms)
    }
}
func JoinRoom(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var req models.JoinRoomRequest

        if err := c.Bind(&req); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
        }

        var room models.Room
        if err := db.First(&room, req.RoomID).Error; err != nil {
            return c.JSON(http.StatusNotFound, map[string]string{"error": "Room not found"})
        }

        if room.RoomType == "private" {
            if req.Code == "" {
                return c.JSON(http.StatusForbidden, map[string]string{"error": "Code is required for private room"})
            }

            if room.RoomCode != req.Code {
                return c.JSON(http.StatusForbidden, map[string]string{"error": "Invalid code for private room"})
            }
        }

        return c.JSON(http.StatusOK, map[string]string{"message": "Successfully joined the room"})
    }
}


