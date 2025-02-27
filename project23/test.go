package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	//c.ShouldBind使用了c.Request.Body，不可重用。
	if errA := c.ShouldBind(&objA); errA == nil {
		c.Sring(http.StatusOK, "the body should be formA")

	}
}
