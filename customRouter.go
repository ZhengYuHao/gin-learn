package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//func main(){
//	router:=gin.Default()
//	http.ListenAndServe(":3333",router)
//}

func main(){
	router:=gin.Default()
	s:=&http.Server{
		Addr:":3333",
		Handler:router,
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,

	}
	s.ListenAndServe()
	router.Run(":3333")
}