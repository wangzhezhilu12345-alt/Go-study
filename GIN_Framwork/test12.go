package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/articles", func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "20"))

		if limit > 100 {
			limit = 100
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []gin.H{}, // articles
			"meta": gin.H{
				"limit":  limit,
				"offset": offset,
				"total":  0, // total
			},
		})
	})
}
