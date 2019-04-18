package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Auth        bool
	HandlerFunc http.HandlerFunc
}
type Routes []Route

//路由配置
var routes = Routes{
	Route{
		"api",
		"GET",
		"/api",
		false,
		IndexHandler,
	},
}
