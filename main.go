package main

import (
	"context"
	"discord-metrics-server/v2/db"
	"discord-metrics-server/v2/messages"
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

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	fmt.Println("Starting server")

	r := gin.Default()
	r.Use(stats.RequestStats())

	r.GET("api/v1/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	fmt.Println("Activating routes")
	messages.Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
