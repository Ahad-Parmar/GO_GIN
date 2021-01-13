package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Ahad Aslam Parmar",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	

	c.JSON(200, gin.H{
		"name": name,
		
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		
	})
}

func main() {
	fmt.Println("Hello world")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)             // /query?name=ahad
	r.GET("/path/:name", PathParameters) // /path/ahad
	r.Run()

}
