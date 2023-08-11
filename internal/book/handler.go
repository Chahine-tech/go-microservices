package book

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	book := router.Group("/book")
	{
		book.GET("/", getHandler)
	}
}

func getHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Check book",
	})
}
