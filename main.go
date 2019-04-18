package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// addr : 接口地址（如“:3000”）
func serve(addr string) {
	Router := NewRouter() //初始化 路由调用

	server := http.Server{
		Addr:         addr,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		Handler:      Router,
	}
	fmt.Println(`🚀 http://localhost` + server.Addr)
	log.Fatal(server.ListenAndServe())
}
func main() {
	serve(`:8080`)
}
