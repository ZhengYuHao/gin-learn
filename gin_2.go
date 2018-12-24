package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

type User struct {

	Username  string ` form:"username" json:"username" binding:"required" `
	Passwd string     `form:"passwd" json:"passwd" bdinding:"required" `
	Age  int         ` form:"age" json:"age" `
 }

func main(){
	router:=gin.Default()
	router.POST("/login",FJSON)
	router.Run(":3333")

}
func FJ(c *gin.Context){
	var user User
	err:=c.Bind(&user)
	if err!=nil{
		fmt.Println(err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"username": user.Username,
		"passwd": user.Passwd,
		"age": user.Age,
	})
}
func FJSON (c * gin.Context){
	var user User
	var err error
	contentType:=c.Request.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		err=c.BindJSON(&user)
	case "application/x-www-form-urlencoded":
		err=c.BindWith(&user,binding.Form)
	}
	if err!=nil{
		fmt.Println(err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"user": user.Username,
		"passwd":user.Passwd,
		"age": user.Age,

	})
}
/*
curl -X POST http://127.0.0.1:3333/login -H "Content-Type:application/x-www-form-urlencoded" -d "username=rsj217&passwd=123&age=21" | python -m json.tool
curl -X POST http://127.0.0.1:3333/login -H "Content-Type:application/x-www-form-urlencoded" -d "username=rsj217&new=21" | python -m json.tool
curl -X POST http://127.0.0.1:3333/login -H "Content-Type:application/json" -d '{"username": "rsj217", "passwd": "123", "age": 21}' | python -m json.tool
 */