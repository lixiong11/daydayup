package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用JSONP向不同域的服务器请求数据，如果查询参数存在回调，则将回调添加到响应体中
func main() {
	r := gin.Default()
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)

	})
	// 监听并在0.0.0.0:8080上启动服务
	r.Run(":8080")
}
