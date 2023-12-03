package main

import (
	"discord-metrics-server/v2/messages"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/semihalev/gin-stats"
	"net/http"
)

func main() {
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
