package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Login struct {
// 	User     string `form:"user" json:"user" xml:"user"  binding:"required"` //这是结构体标签，方便从不同的地方读取用的。
// 	Password string `form:"password" json:"password" xml:"password" binding:"required"`
// }

// func main() {
// 	r := gin.Default()
// 	r.POST("/loginjson", func(c *gin.Context) {
// 		var json Login
// 		if err := c.ShouldBindJSON(&json); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status": "unauthorized",
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
// 	})
// 	r.Run(":8080")
// }
