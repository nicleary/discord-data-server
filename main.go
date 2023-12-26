package main

import (
	"discord-metrics-server/v2/db"
	"discord-metrics-server/v2/messages"
	"discord-metrics-server/v2/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/semihalev/gin-stats"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Creating db context")

	err := db.CreateClient()
	if err != nil {
		log.Fatalf("Error occured while getting client: %v", err)
	}

	client := db.GetClient()
	defer client.Close()

	fmt.Println("Starting server")

	r := gin.Default()
	r.Use(stats.RequestStats())

	r.GET("api/v1/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	fmt.Println("Activating routes")
	messages.Routes(r)
	users.Routes(r)
	r.Run()
}
