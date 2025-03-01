package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 上传文件
func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		//form, _ := c.MultipartForm()
		//files := form.File["upload[]"]
		// for _, file := range files {
		// 	log.Println(file.Filename)

		// 	// Upload the file to specific dst.
		// 	c.SaveUploadedFile(file, dst)
		// }
		//c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

		//若为简单文件代码如下：
		file, _ := c.FormFile("file")
		log.Panicln(file.Filename)
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})
	router.Run(":8080")
}
