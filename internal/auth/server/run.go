package server

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(logger log.Logger, handler http.Handler, port string) {

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("listening on port", port)
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
