package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 静态文件服务
func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resoure/favicon.ico")
	//监听并在0.0.0.0:8080
	router.Run(":8080")
}
