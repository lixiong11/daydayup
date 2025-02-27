package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type formA struct {
	Foo string `josn:"foo" xml:"foo" binding:"required"`
}
type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	//c.ShouldBind使用了c.Request.Body，不可重用。
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, "the body should be formA")
		//因为现在c.Request.Body是EOF，所以这里会报错。
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		//此次省列
	}
}
