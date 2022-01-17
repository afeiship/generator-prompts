package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	r.Handle("DELETE", "/del", func(c *gin.Context) {
		c.String(200, "delete")
	})
	// 对应8路路由
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})
	r.Run()
}
