package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, `welcome!`)
}

// 定义restful 统一返回数据结构
type Result struct {
	ErrCode string      `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

// 统一返回数据结构抽象方法
func result(w http.ResponseWriter, code string, msg string, data interface{}) {
	result := Result{
		ErrCode: code,
		ErrMsg:  msg,
		Data:    data,
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
