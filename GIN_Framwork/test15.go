package main

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // 自定义错误类型，所有的错误用这个定义
// type AppError struct {
// 	Status  int    `json:"-"`
// 	Code    string `json:"code"`
// 	Message string `json:"message"`
// }

// var (
// 	ErrNotFound     = &AppError{Status: 404, Code: "NOT_FOUND", Message: "resource not found"}
// 	ErrUnauthorized = &AppError{Status: 401, Code: "UNAUTHORIZED", Message: "authentication required"}
// 	ErrBadRequest   = &AppError{Status: 400, Code: "BAD_REQUEST", Message: "invalid request"}
// )

// // 这一步太秒了，error这个内置的接口类型，定义了一个简单的方法
// func (e *AppError) Error() string {
// 	return e.Message
// }

// func Errnohandler() gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		c.Next()
// 		if len(c.Errors) == 0 {
// 			return
// 		}
// 		err := c.Errors.Last().Err //取最后一个错误的变量赋值
// 		var apperr *AppError
// 		if errors.As(err, &apperr) {
// 			c.JSON(apperr.Status, gin.H{
// 				"success": false,
// 				"error": gin.H{
// 					"code":    apperr.Code,
// 					"message": apperr.Message},
// 			})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"success": false,
// 				"error":   gin.H{"code": "INTERNAL", "message": "an unexpected eror"},
// 			})

// 		}
// 	}
// }
// func main() {
// 	r := gin.Default()
// 	r.Use(Errnohandler())

// 	r.GET("/api/items/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		if id == "0" {
// 			_ = c.Error(ErrNotFound)
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id}})
// 	})

// 	r.Run(":8080")

// }
