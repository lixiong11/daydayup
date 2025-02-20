package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// 自定义模板渲染器
func main() {
	router := gin.Default()
	html := template.Must(template.ParseFiles("index.tmpl"))
	router.SetHTMLTemplate(html)
	router.Run(":8080")
}

//未执行成功
