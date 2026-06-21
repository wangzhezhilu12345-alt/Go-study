package main

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	router.GET("/cookie", func(c *gin.Context) {
// 		cookie, err := c.Cookie("gin_Cookie")
// 		if err != nil {
// 			cookie = "nonset"
// 			c.SetCookie("gin_Cookie", "test", 3600, "/", "localhost", true, true)

// 		}
// 		fmt.Printf("Cookie values is : %s", cookie)
// 	})
// 	router.Run("localhost:8080")
// }
