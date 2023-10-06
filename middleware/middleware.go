package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	SECRET = "secret"
)

func AuthValid(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token required",
		})
		c.Abort()
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("Invalid Token %s", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if err == nil && token.Valid {
		fmt.Println("Token telah diverifikasi")
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
			"error":   err.Error(),
		})
		c.Abort()
	}
}
