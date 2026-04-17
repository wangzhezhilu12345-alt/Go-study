package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	r.GET("/user/:name/action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")
// 		message := name + " is " + action
// 		c.String(http.StatusOK, message) //写道Http响应体里面。

// 	})

// 	r.Run(":8080")
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/user/:name/*action", func(c *gin.Context) {
// 		firstname:=c.DefaultQuery("firstname","guest")
// 		lastname:=c.Query("lastname")
// 		c.String(http.StatusOK,"Hello %s %s",firstname,lastname)

// 	})

// 	r.Run(":8080")
// }
