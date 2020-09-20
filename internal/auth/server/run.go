package server

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(logger *zap.Logger, handler http.Handler, port string) {

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Info("Listening", zap.String("port", port))
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	}()

	logger.Error("err", zap.Error(<-errs))

}
