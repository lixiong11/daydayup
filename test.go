package main

import (
	// "net/http"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	//1.创建gin对象
	r := gin.Default()
	// //使用LoadHTMLGolb()或者LoadHTMLFile()
	// r.LoadHTMLGlob("templates/**/*")

	// //2. 绑定路由规则，执行的函数
	// //gin.Context,封装了request和response
	// r.GET("/posts/index", func(c *gin.Context) {
	// 	// c.String(http.StatusOK, "hello world!")
	// 	// data := map[string]interface{}{
	// 	// 	"lang": "go语言",
	// 	// 	"tag":  "<br>",
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Posts",
	// 	})

	// })
	// r.GET("users/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "users",
	// 	})
	// })

	// c.AsciiJSON(http.StatusOK, data)
	//3.监听端口，默认在8080
	// getting := func(c *gin.Context) {}
	// r.POST("/xxxpost", getting)
	// html := template.Must(template.ParseFiles("file1", "file2"))
	// r.SetHTMLTemplate(html)
	r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLFiles("./templates/raw.tmpl")
	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})
	r.Run(":8080")

}
