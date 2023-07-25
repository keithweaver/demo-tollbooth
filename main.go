package main

import (
	"demo-tollbooth/cars"
	"demo-tollbooth/middleware"
	"demo-tollbooth/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	MaxRequestsPerSecond = 5 // 5 per second
)

func main() {
	fmt.Println("Starting...")

	carsHandler := cars.NewHandler()

	router := gin.Default()
	router.Use(gin.Recovery())

	carsAPI := router.Group("/cars")
	{
		carsAPI.POST("", user.ValidateAuth(), middleware.LimitHandler(MaxRequestsPerSecond), carsHandler.CreateCar)
	}

	router.Run(":7800")
}
