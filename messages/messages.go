package messages

import (
	"context"
	"discord-metrics-server/v2/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadMessage(c *gin.Context) {
	var message NewDiscordMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("Message", message.MessageData.Contents)
	client := db.GetClient()
	_, err := client.Message.Create().SetContents(message.MessageData.Contents).SetSenderID(message.UserID).SetMessageID(message.MessageData.MessageID).Save(context.Background())

	if err != nil {
		fmt.Println("error creating message object!")
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func Routes(router *gin.Engine) {
	message := router.Group("api/v1/message")
	{
		message.POST("/", UploadMessage)
	}
}
