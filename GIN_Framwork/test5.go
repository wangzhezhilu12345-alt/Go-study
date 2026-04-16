package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file") //这个接口是拿到文件的元信息，文件的内容还没有被读取到内存中。
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		dst := filepath.Join("./files", filepath.Base(file.Filename))
		c.SaveUploadedFile(file, dst)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s uploaded", file.Filename),
		})
	})

	r.Run("localhost:8080")

}
