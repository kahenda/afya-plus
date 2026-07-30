package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/afya-plus/backend/config"
	"github.com/kahenda/afya-plus/backend/handlers"
	"github.com/kahenda/afya-plus/backend/middleware"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	router.Use(middleware.CORSConfig())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Afya Plus backend is running",
		})
	})

	router.POST("/login", handlers.Login)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users", handlers.GetUsers)
	router.POST("/households", handlers.CreateHousehold)
	router.GET("/households", handlers.GetHouseholds)
	router.GET("/households/:id", handlers.GetHouseholdDetail)
	router.POST("/visits", handlers.CreateVisit)
	router.GET("/flags", handlers.GetFlags)
	router.PATCH("/flags/:id", handlers.UpdateFlagStatus)

	router.Run(":8080")
}
