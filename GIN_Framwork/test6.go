package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// const (
// 	maxsize = 1 << 20
// )

// func poster(c *gin.Context) {
// 	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxsize)
// 	err := c.Request.ParseMultipartForm(maxsize)
// 	if err != nil {
// 		if _, ok := err.(*http.MaxBytesError); ok { //这是检查表单里面有没有file这个字段
// 			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
// 				"error": fmt.Sprintf("file too large (max: %d bytes)", maxsize),
// 			})
// 		}
// 		return
// 	}

// 	file, _, err := c.Request.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "No file provided",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "form uploaded successful",
// 	})

// 	file.Close()
// }

// func main() {
// 	r := gin.Default()
// 	r.POST("/upload", poster)
// }
