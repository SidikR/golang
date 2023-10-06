package auth

import (
	"net/http"
	"restAPIs/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	USER     = "admin"
	PASSWORD = "Password123"
	SECRET   = "secret"
)

func LoginHandler(c *gin.Context) {
	var user models.Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return // Return early in case of an error
	}

	if user.Username != USER {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user invalid",
		})
		return // Return early in case of unauthorized user
	}

	if user.Password != PASSWORD {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "password invalid",
		})
		return // Return early in case of invalid password
	}

	// Create JWT token
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		Issuer:    "test",
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return // Return early in case of an error
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   tokenString,
	})
}
