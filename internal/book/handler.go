package book

import (
	"go-microservices/playground"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	book := router.Group("/book")
	{
		book.GET("/", getHandler)
	}
}

func getHandler(c *gin.Context) {
	list := playground.GetBooks()
	err := list.ToJSON(c.Writer)
	if err != nil {
		c.Status(500)
		return
	}
}
