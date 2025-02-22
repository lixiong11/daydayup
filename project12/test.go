package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		//取query里面的参数
		id := c.Query("id")
		//DefaultQuery如果query给值就取给定的值，没有就取默认值0
		page := c.DefaultQuery("page", "0")
		//接收body中form-data里面对应的参数
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id:%s;page:%s;name:%s;message:%s", id, page, name, message)
		c.JSON(200, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})
	router.Run(":8080")
}
