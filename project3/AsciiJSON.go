package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。
func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br>",
		}
		//输出：{"lang":"Go\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
	// 访问地址：http://127.0.0.1:8080/someJSON
	r.Run(":8080")
}
