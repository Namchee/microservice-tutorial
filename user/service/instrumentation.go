package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsMiddleware func(UserService) UserService

type prometheusMiddleware struct {
	requestCount   prometheus.Counter
	requestLatency prometheus.HistogramVec
	next           UserService
}

func NewInstrumentationMiddleware(count prometheus.Counter, latency prometheus.HistogramVec) MetricsMiddleware {
	return func(next UserService) UserService {
		return &prometheusMiddleware{
			requestCount:   count,
			requestLatency: latency,
			next:           next,
		}
	}
}

func (mw *prometheusMiddleware) GetUsers(ctx context.Context, pagination *entity.Pagination) ([]*entity.User, error) {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		mw.requestLatency.WithLabelValues("").Observe(v)
	}))
	defer func() {
		mw.requestCount.Inc()
		timer.ObserveDuration()
	}()
	return mw.next.GetUsers(ctx, pagination)
}

func (mw *prometheusMiddleware) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	return mw.next.GetUserById(ctx, id)
}

func (mw *prometheusMiddleware) CreateUser(ctx context.Context, data *entity.User) (*entity.User, error) {
	return mw.next.CreateUser(ctx, data)
}

func (mw *prometheusMiddleware) DeleteUser(ctx context.Context, id int) (*entity.User, error) {
	return mw.next.DeleteUser(ctx, id)
}
