package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(c *gin.RouterGroup) {
	users := c.Group("/users")
	{
		users.GET("/", listusers)
		users.POST("/", createusers)
	}

}

func listusers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"action": "list_users",
	})
}

func createusers(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"action": "created_users",
	})
}

func RegisterOrderRoutes(c *gin.RouterGroup) {
	orders := c.Group("/orders")
	orders.GET("/:id", getorder)
}

func getorder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"action": "get_order",
	})
}

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")

	RegisterUserRoutes(api)
	RegisterOrderRoutes(api)

	r.Run(":8080")
}
