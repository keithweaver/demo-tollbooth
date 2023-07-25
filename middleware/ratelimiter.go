package middleware

import (
	"demo-tollbooth/auth"
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"time"
)

func LimitHandler(requestsPerSecond float64) gin.HandlerFunc {
	limit := tollbooth.NewLimiter(requestsPerSecond, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	return func(c *gin.Context) {
		session, exists := auth.GetSession(c)
		if !exists {
			c.Abort()
			fmt.Printf("No Auth in Rate limit")
			c.JSON(403, gin.H{})
			return
		}

		httpErr := tollbooth.LimitByKeys(limit, []string{session.ID})
		c.Writer.Header().Add("X-Rate-Limit-Limit", fmt.Sprintf("%.2f", limit.GetMax()))
		c.Writer.Header().Add("X-Rate-Limit-Duration", "1")
		if httpErr != nil {
			c.JSON(429, gin.H{})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
