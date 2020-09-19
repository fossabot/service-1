package logging

import (
	"context"

	"time"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth"
	"github.com/perfolio/service/internal/auth/model"
)

// LoggingMiddleware takes a logger as a dependency
// and returns a service Middleware.
func Use(logger log.Logger) auth.Middleware {
	return func(next auth.Service) auth.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   auth.Service
}

func (mw loggingMiddleware) CreateUser(ctx context.Context, email string, password string) (user model.User, err error) {
	logger := log.With(mw.logger, "method", "CreateUser")

	defer func(begin time.Time) {
		_ = logger.Log(
			"email", email,
			"id", user.ID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	user, err = mw.next.CreateUser(ctx, email, password)
	return
}

func (mw loggingMiddleware) DeleteUser(ctx context.Context, id uuid.UUID) (err error) {
	logger := log.With(mw.logger, "method", "DeleteUser")

	defer func(begin time.Time) {
		_ = logger.Log(
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.next.DeleteUser(ctx, id)
	return
}
