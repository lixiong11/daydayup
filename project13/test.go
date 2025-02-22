package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用SecureJSON防止json劫持，如果给定结构是数组值，则默认预置while(1)到响应体。
func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		name := []string{"lena", "austin", "foot"}
		c.SecureJSON(http.StatusOK, name)
	})
	r.Run(":8080")
}
