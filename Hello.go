package  main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"log"
)

func main(){
	router:=gin.Default()
	//router.POST("/form_post",F)
	//router.POST("/upload",R)
	router.POST("/multi/upload",MutulR)
	router.Run(":3333")
}


func F(c *gin.Context){
	message:=c.PostForm("message")
	nick:=c.DefaultPostForm("nick","anonymous")
	c.JSON(http.StatusOK,gin.H{
		"status":gin.H{
			"status_code":http.StatusOK,
			"status":"ok",
		},
		"message":message,
		"nick":nick,
	})
}

func put(c *gin.Context){
	id:=c.Query("id")
	page:=c.DefaultQuery("page","0")
	name:=c.PostForm("message")
	message:=c.PostForm("message")
	fmt.Printf("id: %s; page: %s; name: %s; message: %s \n",id,page,name,message)
	c.JSON(http.StatusOK,gin.H{
		"status":http.StatusOK,
	})
}


func R(c *gin.Context){//上传一个文件
	name:=c.PostForm("name")
	fmt.Println(name)
	file,header,err:=c.Request.FormFile("upload")
	if err!=nil{
		c.String(http.StatusBadRequest,"Bad request")
		return
	}
	filename:=header.Filename

	fmt.Println(file,err,filename)

	out,err:=os.Create(filename)
	if err!=nil{
		log.Fatal(err)
	}
	defer out.Close()
	_,err=io.Copy(out,file)
	if err!=nil{
		log.Fatal(err)
	}
	c.String(http.StatusCreated,"upload successful")
}
func MutulR(c *gin.Context){
	err:=c.Request.ParseMultipartForm(200000)
	if err!=nil{
		log.Fatal(err)
	}
	formdata:=c.Request.MultipartForm
	files:=formdata.File["upload"]
	for i,_:=range files{
		file,err:=files[i].Open()
		defer file.Close()
		if err!=nil{
			log.Fatal(err)
		}
		out,err:=os.Create(files[i].Filename)

		defer out.Close()

		if err!=nil{
			log.Fatal(err)
		}
		_,err=io.Copy(out,file)

		if err!=nil{
			log.Fatal(err)
		}
		c.String(http.StatusCreated,"upload successful")

	}

}
/*
curl -X POST http://127.0.0.1:3333/upload -F "upload=@/Users/yuhao/Pictures/1.jpg" -H "Content-Type: multipart/form-data"
curl -X POST http://127.0.0.1:3333/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=rsj217" | python -m json.tool
curl -X POST http://127.0.0.1:3333/multi/upload -F "upload=@/Users/yuhao/Pictures/1.jpg" -F "upload=@/Users/yuhao/Pictures/2.jpg" -H "Content-Type: multipart/form-data"

*/
