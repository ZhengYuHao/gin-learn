package  main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
	fmt.Println("before middleware")
	c.Set("request", "clinet_request")
	c.Next()
	fmt.Println("before middleware")
	}
}


func main(){
	router:=gin.Default()
	router.Use(MiddleWare())
	{
		router.GET("/middleware", func(c *gin.Context) {
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			c.JSON(http.StatusOK, gin.H{
				"middile_request": request,
				"request": req,
			})
		})
	}
	router.Run(":3333")
}

