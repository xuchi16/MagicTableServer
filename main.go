package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro.xuchi/magic_table/v2/internal/handlers"
	"pro.xuchi/magic_table/v2/internal/models"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/rooms/create", func(c *gin.Context) {
		var room models.Room
		if err := c.ShouldBindJSON(&room); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Room created successfully",
			"room":    room,
		})
	})

	r.POST("/hall/sendMessage", handlers.SendMessage)
	r.GET("/hall/getMessage", handlers.GetMessage)

	r.GET("/ws", handlers.WebSocketHandle)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
