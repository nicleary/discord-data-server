package messages

import (
	"context"
	"discord-metrics-server/v2/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func UploadMessage(c *gin.Context) {
	var message DiscordMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("Message", message.Contents)
	client := db.GetClient()
	_, err := client.Message.Create().SetContents(message.Contents).Save(context.Background())

	if err != nil {
		fmt.Println("error creating message object!")
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func Routes(router *gin.Engine) {
	message := router.Group("api/v1/message")
	{
		message.GET("/test", Test)
		message.POST("/", UploadMessage)
	}
}
