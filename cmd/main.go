package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	l "github.com/exalor-solution/rest-basic/pkg/xLogger"
)

const (
	logName = "x-app.basic-rest.service"
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
	defer func() {
		logger.Info(ctx, "stopping service...")
		cancel()
		logger.LogSync()
	}()

	logger.Info(ctx, "Loading config...")
	// load config

	srv := http.Server{
		Addr:    "",
		Handler: nil,
	}

	go func() {
		defer log.Println("server has been stopped")
		if err := srv.ListenAndServe(); err != nil {
			internalInterrupt <- err
		}

	}()

	logger.Info(ctx, "service listening for any interrupt signals...")

	select {
	case <-osInterrupt:
		logger.Info(ctx, "OS interrupt signal received")
	case e := <-internalInterrupt:
		logger.Error(ctx, "user service listener interrupt signal received, %+v", zap.Any("lis", e))
	}

}
