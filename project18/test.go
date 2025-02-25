package main

import "github.com/gin-gonic/gin"

type vde struct {
	Bar   string `json:"bar"`
	Apple string `json:"apple"`
}

func ecdg(c *gin.Context) {
	var a vde
	c.BindJSON(&a)
	c.JSON(200, gin.H{
		"bar":   a.Bar,
		"apple": a.Apple,
	})
}

func main() {
	r := gin.Default()
	r.POST("/getting", ecdg)
	r.Run()
}
