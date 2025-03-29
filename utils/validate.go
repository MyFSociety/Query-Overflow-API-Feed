package utils

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	jwt "github.com/golang-jwt/jwt"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func ValidateToken(c *gin.Context) {
	// Get the token from the request header
	token := c.GetHeader("Authorization")

	// Check if the token is empty
	if token == "" {
		c.JSON(401, gin.H{"error": "Authorization token is required"})
		c.Abort()
		return
	}

	// Validate the token

	var valid, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check if the token method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	// Check if the token is valid
	if !valid.Valid {
		c.JSON(401, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	// If the token is valid, proceed to the next handler
	c.Next()
}
