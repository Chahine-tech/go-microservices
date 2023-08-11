package main

import (
	"go-microservices/internal/book"
	"go-microservices/internal/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Check cors",
		})
	})
	book.SetupRoutes(router)
	product.SetupRoutes(router)
	router.Run(":8080")
}
