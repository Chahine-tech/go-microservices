package product

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	product := router.Group("/product")
	{
		product.GET("/", getHandler)
	}
}

func getHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Check product",
	})
}
