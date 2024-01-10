package messages

import (
	"context"
	"discord-metrics-server/v2/db"
	"discord-metrics-server/v2/ent"
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
	messageObject, err := client.
		Message.
		Query().
		Where(message.MessageID(MessageID.MessageID)).
		WithSender().
		WithInReplyTo().
		Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Message ID not found",
		})
	}

	c.JSON(http.StatusOK, MessageToSchema(messageObject))
}

func UploadMessage(c *gin.Context) {
	var incomingMessage DiscordMessage
	if err := c.ShouldBindJSON(&incomingMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Parse datetime field
	timeObject, err := utils.ConvertType(incomingMessage.SentAt)

	if err != nil {
		fmt.Println("Unable to parse datetime")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to time sent",
		})
		return
	}

	client := db.GetClient()

	// Get user object
	userObject, err := client.User.Query().Where(user.UserID(incomingMessage.UserID)).Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User ID for sender not found",
		})
		return
	}

	var inReplyToMessage *ent.Message = nil

	// Get in reply to message object
	if incomingMessage.InReplyTo != "" {
		inReplyToMessage, err = client.
			Message.
			Query().
			Where(message.MessageID(incomingMessage.InReplyTo)).
			Only(context.Background())
		if err != nil {
			fmt.Println("Invalid in reply to message")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid in_reply_to message ID",
			})
		}
	}
	fmt.Println("Here it is:")

	var messageObject *ent.Message = nil

	if inReplyToMessage != nil {
		messageObject, err = client.Message.Create().
			SetContents(incomingMessage.Contents).
			SetSender(userObject).
			SetMessageID(incomingMessage.MessageID).
			SetSentAt(timeObject).
			SetInReplyTo(inReplyToMessage).
			SetChannelID(incomingMessage.ChannelID).
			Save(context.Background())
	} else {
		messageObject, err = client.Message.Create().
			SetContents(incomingMessage.Contents).
			SetSender(userObject).
			SetMessageID(incomingMessage.MessageID).
			SetSentAt(timeObject).
			SetChannelID(incomingMessage.ChannelID).
			Save(context.Background())
	}

	if err != nil {
		fmt.Println("error creating incomingMessage object!")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid incomingMessage object",
		})
		return
	}
	messageObject, _ = client.Message.
		Query().
		Where(message.MessageID(messageObject.MessageID)).
		WithSender().
		WithInReplyTo().
		First(context.Background())
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

	messagesQuery := client.Message.Query()

	if MessageQuery.UserID != "" {
		userObject, err := client.User.Query().Where(user.UserID(MessageQuery.UserID)).Only(context.Background())
		if err != nil {
			fmt.Println("Error getting user")
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID",
			})
		}
		messagesQuery.Where(message.SenderID(userObject.ID))
	}

	if MessageQuery.ChannelID != "" {
		messagesQuery.Where(message.ChannelID(MessageQuery.ChannelID))
	}

	messages, mesErr := messagesQuery.
		Offset(offset).
		Limit(MessageQuery.PageSize).
		WithSender().
		WithInReplyTo().
		All(context.Background())

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
