package user

import (
	"demo-tollbooth/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ValidateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.AbortWithStatus(403)
			return
		}

		authToken = strings.ReplaceAll(authToken, "Bearer ", "")

		session := auth.Session{
			ID: authToken,
		}
		// TODO - Verify session is valid, THIS IS BEING SKIPPED
		if authToken == "" { // This is just a demo use case
			fmt.Println("Session expired")
			c.AbortWithStatus(403)
			return
		}
		c.Set("session", session)

		c.Next()
	}
}
