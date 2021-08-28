package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"short_url/lib/mux"
)

var (
	port string
)

func main() {

	// 获取命令行 启动端口号
	flag.StringVar(&port, "port", "8080", "http port")
	// 解析端口
	flag.Parse()
	//
	log.Println("web server start at port ", port)
	//http.Handle("/hello/golang/",&BaseHandler{})
	//http.HandleFunc("/hello/world", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello world"))
	//})
	//http.ListenAndServe(":"+port, nil)
	router := mux.NewMuxHandler()

	router.Handler("/hello/golang/",&BaseHandler{})
	router.HandleFunc("/hello/world", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello world"))
	})

	log.Fatalln(http.ListenAndServe(":8080",router))

}

type BaseHandler struct {
}

func (handler *BaseHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path => ", req.URL.Path)
	fmt.Println("url param a=>", req.URL.Query().Get("a"))
	resp.Write([]byte("hello golang"))
}
