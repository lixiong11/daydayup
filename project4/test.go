package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由对象
	router := gin.Default()
	//使用LoadHTMLGlob()或者LoadHTMLFiles()
	// templates目录下的index.tmpl文件
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			//访问提示页面
			"title": "Main website",
		})
	})
	//访问地址：http://127.0.0.1:8080/index
	router.Run(":8080")
}
