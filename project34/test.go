package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.Any("testing", startPage)
	router.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====Only Bind By Query String")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

//http://127.0.0.1:8085/testing?name=xiaoming&address=hubei
//person.Name: xiaoming   person.Address= hubei
//第16行any 为任意请求方式
