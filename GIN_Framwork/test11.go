package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func fail(c *gin.Context, code int, errCode string, message string) {
	c.JSON(code, gin.H{
		"Success": false,
		"error": gin.H{
			"Code":    errCode,
			"Message": message,
		},
	})
}

func OK(c *gin.Context, data interface{}) {
	//函数设计的接口类型是为了任何值都能传进来。
	c.JSON(200, gin.H{
		"data":    data,
		"success": true,
	})
}

func main() {
	router := gin.Default()

	router.GET("api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "0" {
			fail(c, http.StatusNotFound, "USER_NOT_FOUND", "no user with that ID")
		}
		OK(c, gin.H{"id": id, "name": "Alice"})
	})

}
