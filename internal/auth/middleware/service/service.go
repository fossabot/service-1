package service

import (
	"context"

	"github.com/perfolio/service/internal/auth"
	"github.com/perfolio/service/internal/auth/model"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware takes a logger as a dependency
// and returns a service Middleware.
func Logging(logger log.Logger) Middleware {
	return func(next auth.Service) auth.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   auth.Service
}

func (mw loggingMiddleware) CreateUser(ctx context.Context, email string, password string) (user model.User, err error) {
	defer func() {
		mw.logger.Log("method", "CreateUser", "email", email, "id", user.ID, "err", err)
	}()
	user, err = mw.next.CreateUser(ctx, email, password)
	return user, err
}
