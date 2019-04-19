// package main

// import (
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/rs/cors"
// )

// func NewRouter() *mux.Router {
// 	// cors 配置
// 	c := cors.New(cors.Options{
// 		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
// 		AllowCredentials: true,
// 	})
// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routes {
// 		// cors 实现
// 		handler := c.Handler(route.HandlerFunc)

// 		router.
// 			Methods(route.Method...).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(handler)
// 	}
// 	return router
// }

package main

import (
	"github.com/julienschmidt/httprouter"
)

//Reads from the routes slice to translate the values to httprouter.Handle
func NewRouter(routes Routes) *httprouter.Router {

	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc
		handle = Logger(handle)

		router.Handle(route.Method, route.Path, handle)
	}

	return router
}
