package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/afya-plus/backend/config"
	"github.com/kahenda/afya-plus/backend/handlers"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Afya Plus backend is running",
		})
	})

	router.POST("/login", handlers.Login)
	router.POST("/users", handlers.CreateUser)
	router.POST("/households", handlers.CreateHousehold)
	router.POST("/visits", handlers.CreateVisit)

	router.Run(":8080")
}
