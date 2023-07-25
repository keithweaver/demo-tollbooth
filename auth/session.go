package auth

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Session struct {
	ID        string    `json:"id"`
	UserEmail string    `json:"userEmail"`
	ExpiryAt  time.Time `json:"expiry_at"`
}

func GetSession(c *gin.Context) (Session, bool) {
	i, exists := c.Get("session")
	if !exists {
		return Session{}, false
	}
	session, ok := i.(Session)
	if !ok {
		return Session{}, false
	}
	return session, true
}
