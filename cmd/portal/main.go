package main

import (
	"htmx-go/internals/db"
	"htmx-go/internals/models"
	"htmx-go/internals/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := db.Connect()

	err := db.Migrate(conn, &models.User{}, &models.Order{})
	if err != nil {
		panic("Failed to migrate DB.")
	}

	r := server.SetupServer()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}
