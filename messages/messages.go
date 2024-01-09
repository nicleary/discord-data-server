package messages

import (
	"context"
	"discord-metrics-server/v2/db"
	"discord-metrics-server/v2/ent/message"
	"discord-metrics-server/v2/ent/user"
	"discord-metrics-server/v2/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessage(c *gin.Context) {
	var MessageID DiscordMessageID
	if err := c.ShouldBindUri(&MessageID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	client := db.GetClient()
	messageObject, err := client.Message.Query().Where(message.MessageID(MessageID.MessageID)).Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Message ID not found",
		})
	}

	c.JSON(http.StatusOK, MessageToSchema(messageObject))
}

func UploadMessage(c *gin.Context) {
	var message DiscordMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Parse datetime field
	timeObject, err := utils.ConvertType(message.SentAt)

	if err != nil {
		fmt.Println("Unable to parse datetime")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to time sent",
		})
		return
	}

	client := db.GetClient()

	// Get user object
	userObject, err := client.User.Query().Where(user.UserID(message.UserID)).Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User ID for sender not found",
		})
		return
	}

	messageObject, err := client.Message.Create().
		SetContents(message.Contents).
		SetSender(userObject).
		SetMessageID(message.MessageID).
		SetSentAt(timeObject).
		Save(context.Background())

	if err != nil {
		fmt.Println("error creating message object!")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid message object",
		})
		return
	}

	c.JSON(http.StatusOK, MessageToSchema(messageObject))
}

func GetMessages(c *gin.Context) {
	var MessageQuery DiscordMessageQuery
	if c.ShouldBind(&MessageQuery) != nil {
		fmt.Println("Error binding query string")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query string",
		})
		return
	}

	if MessageQuery.PageNumber == 0 {
		MessageQuery.PageNumber = 1
	}
	if MessageQuery.PageSize == 0 {
		MessageQuery.PageSize = 20
	}

	err := MessageQuery.validate()

	if err != nil {
		fmt.Println("Invalid query string")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	client := db.GetClient()

	offset := (MessageQuery.PageNumber - 1) * MessageQuery.PageSize

	messages, mesErr := client.Message.Query().Offset(offset).Limit(MessageQuery.PageSize).All(context.Background())

	if mesErr != nil {
		fmt.Println("Error pulling messages from DB")
		fmt.Println(mesErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while pulling messages",
		})
		return
	}

	var messageJSONs []DiscordMessageResponse

	for _, messageObject := range messages {
		messageJSONs = append(messageJSONs, MessageToSchema(messageObject))
	}

	c.JSON(http.StatusOK, messageJSONs)
}

func Routes(router *gin.Engine) {
	message := router.Group("api/v1/message")
	{
		message.GET("/", GetMessages)
		message.GET("/:id", GetMessage)
		message.POST("/", UploadMessage)
	}
}
