package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/exalor-solution/rest-basic/transport/xHttp"
	"go.uber.org/zap"

	l "github.com/exalor-solution/rest-basic/pkg/xLogger"
)

const (
	logName = "x-app.basic-rest.service"
	address = "0.0.0.0:8080"
)

var (
	osInterrupt       chan os.Signal
	internalInterrupt chan error
)

func init() {
	osInterrupt = make(chan os.Signal)
	internalInterrupt = make(chan error)
	signal.Notify(osInterrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := l.NewLogger(logName)

	srv := http.Server{
		Addr:    address,
		Handler: xHttp.Run(ctx, logger),
	}
	defer func() {
		logger.Info(ctx, "stopping service...")
		_ = srv.Shutdown(ctx)
		cancel()
		logger.LogSync()
	}()

	go func() {
		defer log.Println("server has been stopped")
		logger.Info(ctx, "server has been started at ", zap.String("address", address))
		if err := srv.ListenAndServe(); err != nil {
			internalInterrupt <- err
		}
	}()

	logger.Info(ctx, "service listening for any interrupt signals...")

	select {
	case <-osInterrupt:
		logger.Info(ctx, "OS interrupt signal received")
	case e := <-internalInterrupt:
		logger.Error(ctx, "service listener interrupt, %+v", zap.Error(e))
	}

}
