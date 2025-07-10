package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main(){
	router := gin.Default()

	router.GET("/",func (c *gin.Context){
		c.JSON(200,gin.H{
			"msg":"Hello World!",
		})
	})

	router.Use(cors.Default())
	router.Run(":3000")
}