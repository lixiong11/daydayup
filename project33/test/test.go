package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		//map
		//字典 key value
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		fmt.Printf("ids:%v;names:%v,%s", ids, names, "1111")
		c.JSON(200, nil)
		//请求地址：http://127.0.0.1:8080/post  query参数：ids[a]=123，ids[b]=hello;form-data参数：names[aaa]=xiaoming，names[bbb]=bbb
		// ids:map[a:1234 b:hello];names:map[aaa:xiaoming bbb:bbb],1111
		//ids: map[b:hello a:1234], names: map[second:tianou first:thinkerou]
	})
	router.Run(":8080")
}
