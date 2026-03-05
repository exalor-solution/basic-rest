package xHttp

import (
	"io"
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
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		writer.WriteHeader(http.StatusNotImplemented)
		_, _ = writer.Write([]byte("/ is not implemented, "))
	}

	dic[add] = func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if request.Method != "POST" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = writer.Write([]byte("/add is not allowed, "))
			return
		}
		byt, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}

		err = srv.Add(byt)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}

	}

	dic[del] = func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if request.Method != "DELETE" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
	dic[update] = func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if request.Method != "PUT" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
	dic[find] = func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if request.Method != "GET" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}
