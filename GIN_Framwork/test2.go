package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET",
	})
}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}
func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT",
	})
}
func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE",
	})
}

func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func head(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}
func main() {
	router := gin.Default()
	router.GET("/somget", getting)
	router.POST("/sompost", posting)
	router.PUT("/somput", putting)
	router.DELETE("/somdel", deleting)
	router.PATCH("/sompatch", patching)
	router.HEAD("/somhead", head)
	router.OPTIONS("/somoptions", options)
	router.Run()
}
