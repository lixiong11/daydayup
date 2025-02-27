package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 使用现有的基础请求对象解析查询字符串参数。
	// 示例 URL： /welcome?firstname=Jane&lastname=Doe
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}
