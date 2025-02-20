package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("http").Parse(`
<html>
<head>
	<title>Http Test</title>
</head>
<body>
	<h1 style="color:red;">Welcome,Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)
	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			//使用pusher.Push()做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push:%v", err)
			}
		}
		c.HTML(200, "http", gin.H{
			"status": "success",
		})
	})
	r.Run(":8080")
}
