package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth"
	"github.com/perfolio/service/internal/auth/model"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type metricsMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           auth.Service
}

func Use() auth.Middleware {
	return func(next auth.Service) auth.Service {

		fieldKeys := []string{"method", "error"}
		requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "perfolio",
			Subsystem: "auth",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys)
		requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "perfolio",
			Subsystem: "auth",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys)

		return metricsMiddleware{requestCount, requestLatency, next}

	}

}

func (mw metricsMiddleware) CreateUser(ctx context.Context, email string, password string) (user model.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateUser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	user, err = mw.next.CreateUser(ctx, email, password)
	return
}

func (mw metricsMiddleware) DeleteUser(ctx context.Context, id uuid.UUID) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DeleteUser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = mw.next.DeleteUser(ctx, id)
	return
}
