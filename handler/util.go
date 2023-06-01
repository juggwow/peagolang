package handler

import (
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func sendError(c *gin.Context, status int, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"message": err.Error(),
	})
}

func verifyToken(token string) (*claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
