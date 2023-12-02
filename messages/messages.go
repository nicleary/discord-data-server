package messages

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func Routes(router *gin.Engine) {
	message := router.Group("/message")
	{
		message.GET("/test", Test)
	}
}
