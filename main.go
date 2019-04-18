package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	port = "10085" //配置端口
)

func main() {
	Router := NewRouter() //初始化 路由调用

	server := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		Handler:      Router,
	}
	fmt.Println(`🚀 http://localhost:` + port)
	log.Fatal(server.ListenAndServe())
}
