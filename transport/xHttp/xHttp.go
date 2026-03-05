package xHttp

import (
	"net/http"
)

func Run() http.Handler {
	load()
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if f, ok := dic[req.URL.Path]; ok {
			f(w, req)
			return
		}
		http.NotFound(w, req)
	})
}
