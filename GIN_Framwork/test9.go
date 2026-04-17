package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	router := gin.Default()
// 	api := router.Group("/api")
// 	{
// 		v1 := api.Group("/v1")
// 		{
// 			users := v1.Group("/users")
// 			users.GET("/", func(c *gin.Context) {
// 				c.JSON(200, gin.H{
// 					"message": "v1 users",
// 				})
// 			})
// 			users.GET("/:id", func(c *gin.Context) {
// 				c.JSON(200, gin.H{
// 					"message": "v1 user " + c.Param("id"),
// 				})
// 			})

// 			id := v1.Group("/id")
// 		}
// 	}

// 	router.Run(":8080")
// }
