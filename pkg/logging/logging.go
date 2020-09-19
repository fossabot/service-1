package logging

import (
	"github.com/go-kit/kit/log"
	"os"
)

func New(serviceName string) log.Logger {

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", serviceName,
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
	return logger
}
