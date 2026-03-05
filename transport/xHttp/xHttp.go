package xHttp

import (
	"context"
	"net/http"

	"github.com/exalor-solution/rest-basic/pkg/xLogger"
)

func Run(ctx context.Context, l xLogger.ILogger) http.Handler {
	load(ctx, l)
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if f, ok := dic[req.URL.Path]; ok {
			f(w, req)
			return
		}
		http.NotFound(w, req)
	})
}
