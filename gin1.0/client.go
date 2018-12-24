package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func helpRead(resp *http.Response){
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("ERROR!:",err)
	}
	fmt.Println(string(body))
}
func main() {
	//调用最基本的GET,并获得返回值
	resp, _ := http.Get("http://0.0.0.0:3333/test1")
	helpRead(resp)
	//调用最简单的POST,并获得返回值
	resp, _ = http.Post("http://0.0.0.0:3333/test2", "", strings.NewReader(""))
	helpRead(resp)
}