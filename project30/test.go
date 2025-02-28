// 使用 BasicAuth中间件
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123332"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

// 使用BasicAuth中间件
func main() {
	r := gin.Default()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello",
		"manu":   "4321",
	}))
	authorized.GET("/secrets", func(c *gin.Context) {
		//get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET:()"})
		}
	})
	r.Run(":8080")
}
