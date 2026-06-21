package main

// import (
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/sync/errgroup"
// )

// var (
// 	g errgroup.Group
// )

// func router01() http.Handler {
// 	e := gin.New()
// 	e.Use(gin.Recovery())
// 	e.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    http.StatusOK,
// 			"message": "Welcome server 01",
// 		})
// 	})
// 	return e
// }

// func router02() http.Handler {
// 	e := gin.New()
// 	e.Use(gin.Recovery())
// 	e.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    http.StatusOK,
// 			"message": "Welcome server 02",
// 		})
// 	})

// 	return e
// }

// func main() {
// 	gin.SetMode(gin.ReleaseMode)

// 	server01 := &http.Server{
// 		Addr:         "localhost:8080",
// 		Handler:      router01(),
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 	}

// 	server02 := &http.Server{
// 		Addr:         "localhost:8081",
// 		Handler:      router02(),
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 	}

// 	g.Go(func() error {
// 		return server01.ListenAndServe()
// 	})

// 	g.Go(func() error {
// 		return server02.ListenAndServe()
// 	})

// 	if err := g.Wait(); err != nil {
// 		panic(err)
// 	}
// 	//主goroutine需要使用wait等待子goroutine完成，不然提前结束有风险

// }
