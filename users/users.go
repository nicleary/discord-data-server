package users

import (
	"context"
	"discord-metrics-server/v2/db"
	"discord-metrics-server/v2/ent/user"
	"discord-metrics-server/v2/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadUser(c *gin.Context) {
	var user NewDiscordUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Parse datetime field
	timeObject, err := utils.ConvertType(user.DateJoined)

	if err != nil {
		fmt.Println("Unable to parse datetime")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse date joined",
		})
		return
	}

	client := db.GetClient()
	userObject, err := client.User.Create().
		SetUserID(user.UserID).
		SetDateJoined(timeObject).
		SetIsBot(user.IsBot).
		Save(context.Background())
	if err != nil {
		fmt.Println("error creating user object")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})
		return
	}
	c.JSON(http.StatusCreated, UserToSchema(userObject))
}

func GetUser(c *gin.Context) {
	var UserID DiscordUserID
	if err := c.ShouldBindUri(&UserID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	client := db.GetClient()
	userObject, err := client.User.Query().Where(user.UserID(UserID.UserID)).Only(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User ID not found",
		})
		return
	}
	c.JSON(http.StatusOK, UserToSchema(userObject))
}

func Routes(router *gin.Engine) {
	user := router.Group("api/v1/user")
	{
		user.POST("/", UploadUser)
		user.GET("/:id", GetUser)
	}
}
