package cars

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) CreateCar(c *gin.Context) {
	time.Sleep(1 * time.Second) // Simulate work

	c.JSON(201, gin.H{"message": "Car created"})
}
