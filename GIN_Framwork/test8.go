package main

// import "github.com/gin-gonic/gin"

// func distant() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//检查token,session,cookie等
// 		c.Next()
// 	}
// }
// func main() {
// 	router := gin.Default()
// 	public := router.Group("/api")
// 	{
// 		public.GET("/health")
// 	}

// 	private := router.Group("/api")
// 	private.Use(distant())
// 	{
// 		private.GET("/user")
// 	}
// 	router.Run(":8080")
// }
