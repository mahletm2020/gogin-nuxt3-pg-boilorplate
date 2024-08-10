package main

import (
	// "fmt"
	"boilermodule/initilazers"
	"github.com/gin-gonic/gin"
)

func init(){
	// initilazers.LoadEnvVar()
	initilazers.LoadEnvVar()
}
func main(){

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()}
