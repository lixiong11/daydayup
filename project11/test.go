package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 提供unicode实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	//监听并在0.0.0.0:8080上启动服务
	r.Run(":8080")
}
