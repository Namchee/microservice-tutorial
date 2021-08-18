package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsMiddleware func(UserService) UserService

type prometheusMiddleware struct {
	requestCount   prometheus.CounterVec
	requestLatency prometheus.HistogramVec
	next           UserService
}

func NewInstrumentationMiddleware(count prometheus.CounterVec, latency prometheus.HistogramVec) MetricsMiddleware {
	return func(next UserService) UserService {
		return &prometheusMiddleware{
			requestCount:   count,
			requestLatency: latency,
			next:           next,
		}
	}
}

func (mw *prometheusMiddleware) GetUsers(ctx context.Context, pagination *entity.Pagination) ([]*entity.User, error) {
	timer := prometheus.NewTimer(mw.requestLatency.WithLabelValues("get_users"))
	defer func() {
		mw.requestCount.WithLabelValues("get_users").Inc()
		timer.ObserveDuration()
	}()
	return mw.next.GetUsers(ctx, pagination)
}

func (mw *prometheusMiddleware) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	timer := prometheus.NewTimer(mw.requestLatency.WithLabelValues("get_user_by_id"))
	defer func() {
		mw.requestCount.WithLabelValues("get_user_by_id").Inc()
		timer.ObserveDuration()
	}()
	return mw.next.GetUserById(ctx, id)
}

func (mw *prometheusMiddleware) CreateUser(ctx context.Context, data *entity.User) (*entity.User, error) {
	timer := prometheus.NewTimer(mw.requestLatency.WithLabelValues("create_user"))
	defer func() {
		mw.requestCount.WithLabelValues("create_user").Inc()
		timer.ObserveDuration()
	}()
	return mw.next.CreateUser(ctx, data)
}

func (mw *prometheusMiddleware) DeleteUser(ctx context.Context, id int) (*entity.User, error) {
	timer := prometheus.NewTimer(mw.requestLatency.WithLabelValues("delete_user"))
	defer func() {
		mw.requestCount.WithLabelValues("delete_user").Inc()
		timer.ObserveDuration()
	}()
	return mw.next.DeleteUser(ctx, id)
}
