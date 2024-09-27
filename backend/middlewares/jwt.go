package middlewares

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(userID uint) (string, error) {
    claims := Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), 
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET"))) 
}

func VerifyJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil 
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    return claims, nil
}