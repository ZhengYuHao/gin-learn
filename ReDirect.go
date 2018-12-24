package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router:=gin.Default()
	router.GET("/redict/google",func (c *gin.Context){
		c.Redirect(http.StatusMovedPermanently,"https://baidu.com")
	})
	router.Run(":3333")
}