package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	// Serve a file inline (displayed in browser)
// 	router.GET("/local/file", func(c *gin.Context) {
// 		c.File("local/file.go")
// 	})

// 	// Serve a file from an http.FileSystem
// 	var fs http.FileSystem = http.Dir("/var/www/assets")
// 	router.GET("/fs/file", func(c *gin.Context) {
// 		c.FileFromFS("fs/file.go", fs)
// 	})

// 	// Serve a file as a downloadable attachment with a custom filename
// 	router.GET("/download", func(c *gin.Context) {
// 		c.FileAttachment("local/report-2024-q1.xlsx", "quarterly-report.xlsx")
// 	})

// 	router.Run(":8080")
// }
