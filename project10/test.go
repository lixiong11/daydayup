package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/form_post", func(c *gin.Context) {
		//c.PostForm("message"):从POSt请求的表单数据中获取message的字段值如果字段不存在返回为空
		message := c.PostForm("message")
		//c.DefaultPostForm("nick", "tony")：从POST请求的表单数中获取nick的值，若字段不存在返回空字符串
		nick := c.DefaultPostForm("nick", "tony")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
