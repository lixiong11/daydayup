package main

import "github.com/gin-gonic/gin"

func main() {
	//创建gin对象
	r := gin.Default()
	// GET设置url访问路径，gin.Context封装了request与response方法
	r.GET("/ping", func(c *gin.Context) {
		//以json的格式显示访问内容，200为访问成功的状态码
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//监听默认端口 0.0.0.0:8080,浏览器访问路径：http://127.0.0.1:8080/ping
	r.Run()
}
