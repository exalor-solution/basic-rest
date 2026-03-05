package xHttp

import (
	"net/http"

	"github.com/exalor-solution/rest-basic/pkg/service"
)

const (
	root   = "/"
	add    = "/add"
	del    = "/del"
	update = "/update"
	find   = "/find"
)

var (
	dic map[string]func(http.ResponseWriter, *http.Request)
	srv service.ISubscription
)

func init() {
	dic = make(map[string]func(http.ResponseWriter, *http.Request), 5)
	srv = service.New()

}

func load() {

	dic[root] = func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		writer.WriteHeader(http.StatusNotImplemented)
		_, _ = writer.Write([]byte("/ is not implemented, "))
	}

	dic[add] = func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = writer.Write([]byte("/add is not allowed, "))
			return
		}
		err := srv.Add()
		if err != nil {

		}
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	dic[del] = func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "DELETE" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
	dic[update] = func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "PUT" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
	dic[find] = func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
