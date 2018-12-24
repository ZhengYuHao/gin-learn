package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func func1(c *gin.Context){
	c.String(http.StatusOK,"test1 ok")
}
func func2(c *gin.Context){
	c.String(http.StatusOK,"test2 ok")

}
func main(){
	router:=gin.Default()
	router.GET("/test1",func1)
	router.POST("/test2",func2)
	router.Run(":3333")

}
