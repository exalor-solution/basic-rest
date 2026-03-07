package xHttp

import (
	"context"
	"io"
	"net/http"

	"github.com/exalor-solution/rest-basic/pkg/service"
	"github.com/exalor-solution/rest-basic/pkg/xLogger"
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
}

func load(ctx context.Context, l xLogger.ILogger) {
	srv = service.New(l)

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
		byt, code := read(request.Body)
		if code != http.StatusOK {
			writer.WriteHeader(http.StatusInternalServerError)
		}

		res := srv.Add(ctx, byt)

		writer.WriteHeader(res.HttpStatus)
		_, _ = writer.Write([]byte(res.Error()))

	}

	dic[del] = func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if request.Method != "DELETE" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		res := srv.Delete(ctx, request.URL.Query().Get("name"))

		writer.WriteHeader(res.HttpStatus)
		_, _ = writer.Write([]byte(res.Error()))

	}
	dic[update] = func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if request.Method != "PUT" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		byt, code := read(request.Body)
		if code != http.StatusOK {
			writer.WriteHeader(http.StatusInternalServerError)
		}

		res := srv.Update(ctx, request.URL.Query().Get("name"), byt)

		writer.WriteHeader(res.HttpStatus)
		_, _ = writer.Write([]byte(res.Error()))
	}
	dic[find] = func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if request.Method != "GET" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		res, err := srv.Find(ctx, request.URL.Query().Get("name"))
		if err.HttpStatus != 200 {
			writer.WriteHeader(err.HttpStatus)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}

		writer.WriteHeader(err.HttpStatus)
		_, _ = writer.Write([]byte(res))
	}
}

func read(reader io.Reader) ([]byte, int) {
	byt, err := io.ReadAll(reader)
	if err != nil {
		return nil, http.StatusMethodNotAllowed
	}
	return byt, http.StatusOK
}
