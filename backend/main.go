package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/afya-plus/backend/config"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Afya Plus backend is running",
		})
	})

	router.Run(":8080")
}
