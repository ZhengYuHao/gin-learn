package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router:=gin.Default()
	router.GET("/render",func( c *gin.Context){
		contentType:=c.DefaultQuery("content_type","json")
		if contentType=="json"{
			c.JSON(http.StatusOK,gin.H{
				"user": "rsj217",
				"passwd":"123",
			})
		}else if contentType=="xml"{
			c.XML(http.StatusOK,gin.H{
				"user": "rsj217",
				"passwd": "123",
			})
		}

	})
	router.Run(":3333")
}
