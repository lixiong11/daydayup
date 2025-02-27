package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//默认路由日志格式：
//[GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
//[GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
//[GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
//`如果你想要以指定的格式（例如 JSON，key values 或其他格式）记录信息，则可以使用 gin.DebugPrintRouteFunc 指定格式。 在下面的示例中，我们使用标准日志包记录所有路由，但你可以使用其他满足你需求的日志工具。`

func main() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandler int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandler)
	}
	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})
	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	//监听并在0.0.0.0:8080
	r.Run(":8080")
}
