package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	// binding:"required"代表参数验证，必须json传参有该参数
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 你可以使用显示绑定声明绑定 multipart form:
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用ShouldBind方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		err := c.ShouldBind(&form)
		if err == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		} else {
			//log.Panic(err)
			c.JSON(500, gin.H{"status": err.Error()})
		}
	})
	router.Run(":8080")
}
