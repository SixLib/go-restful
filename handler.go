package main

import (
	"encoding/json"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(`welcome!`); err != nil {
		panic(err)
	}
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
