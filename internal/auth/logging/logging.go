package logging

import (
	"context"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth"
	"github.com/perfolio/service/internal/auth/model"
	"go.uber.org/zap"
	"time"
)

// LoggingMiddleware takes a logger as a dependency
// and returns a service Middleware.
func Use(logger *zap.Logger) auth.Middleware {
	return func(next auth.Service) auth.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger *zap.Logger
	next   auth.Service
}

func (mw loggingMiddleware) CreateUser(ctx context.Context, email string, password string) (user model.User, err error) {
	logger := mw.logger.With(zap.String("method", "CreateUser"))

	defer func(begin time.Time) {
		logger.Info(
			"Created new user",
			zap.String("email", email),
			zap.String("id", user.ID.String()),
			zap.Error(err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())
	user, err = mw.next.CreateUser(ctx, email, password)
	return
}

func (mw loggingMiddleware) DeleteUser(ctx context.Context, id uuid.UUID) (err error) {
	logger := mw.logger.With(zap.String("method", "DeleteUser"))

	defer func(begin time.Time) {
		logger.Info(
			"Deleted user",
			zap.String("id", id.String()),
			zap.Error(err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())
	err = mw.next.DeleteUser(ctx, id)
	return
}

func (mw loggingMiddleware) ChangeEmail(ctx context.Context, id uuid.UUID, newEmail string) (err error) {
	logger := mw.logger.With(zap.String("method", "ChangeEmail"))

	defer func(begin time.Time) {
		logger.Info(
			"changed email",
			zap.String("id", id.String()),
			zap.Error(err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())
	err = mw.next.ChangeEmail(ctx, id, newEmail)
	return
}

func (mw loggingMiddleware) ConfirmEmail(ctx context.Context, id uuid.UUID) (err error) {
	logger := mw.logger.With(zap.String("method", "ConfirmEmail"))

	defer func(begin time.Time) {
		logger.Info(
			"changed email",
			zap.String("id", id.String()),
			zap.Error(err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())
	err = mw.next.ConfirmEmail(ctx, id)
	return
}
