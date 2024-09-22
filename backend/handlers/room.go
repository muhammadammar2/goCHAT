// package handlers

// import (
// 	"dummy/models"
// 	"net/http"
// 	"time"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/labstack/echo/v4"

// 	"gorm.io/gorm"
// )

// func CreateRoom(db *gorm.DB) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         user := c.Get("user").(*jwt.Token)
//         claims := user.Claims.(jwt.MapClaims)
//         ownerID := uint(claims["userID"].(float64))

//         var room struct {
//             Name        string `json:"name" validate:"required"`
//             Description string `json:"description"`
//             Type        string `json:"type" validate:"required,oneof=public private"`
//             Code        string `json:"code,omitempty"`
//         }

//         if err := c.Bind(&room); err != nil {
//             return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
//         }

//         // new room
//         newRoom := models.Room{
//             Name:        room.Name,
//             Description: room.Description,
//             Type:        room.Type,
//             OwnerID:     ownerID,
//             CreatedAt:   time.Now(),
//             UpdatedAt:   time.Now(),
//         }

//         if err := db.Create(&newRoom).Error; err != nil {
//             return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Couldn't create room"})
//         }

//         return c.JSON(http.StatusOK, echo.Map{
//             "room": newRoom,
//         })
//     }
// }

package handlers

import (
	"dummy/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateRoom(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Get the Authorization header
        tokenString := c.Request().Header.Get("Authorization")

        // If the Authorization header is empty or improperly formatted
        if tokenString == "" || !startsWithBearer(tokenString) {
            log.Println("Bearer token not found or improperly formatted")
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
        }

        // Remove the "Bearer " prefix from the token string
        tokenString = tokenString[7:] // Strip out the "Bearer " part
		log.Println("Token before parsing:", tokenString)


        // Parse the JWT and extract user claims (error if parsing fails)
        userID, err := GetUserIDFromToken(tokenString)
        if err != nil || userID == 0 {	
            log.Println("Invalid token or failed to retrieve userID")
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
        }

        // Continue with room creation...
        room := new(models.Room)
        if err := c.Bind(room); err != nil {
            log.Println("Failed to bind request body to Room struct")
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
        }

        // Cast userID to uint for room.OwnerID
        room.OwnerID = uint(userID)

        // Create the room in the database
        if err := db.Create(&room).Error; err != nil {
            log.Println("Error creating room:", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create room"})
        }

        return c.JSON(http.StatusOK, room)
    }
}

// Helper function to check if the Authorization header starts with "Bearer"
func startsWithBearer(token string) bool {
    return len(token) > 7 && token[:7] == "Bearer "
}


// Dummy function to simulate token parsing
func GetUserIDFromToken(token string) (int, error) {
	// Simulate extracting user ID from token
	// In reality, you would parse the token and extract the claims
	return 49, nil // Returning 49 as a placeholder userID
}
