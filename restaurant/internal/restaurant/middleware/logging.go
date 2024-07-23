package middleware

import (
	"restaurant/internal/restaurant/domain"
	"restaurant/internal/restaurant/service"
	"time"

	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

func NewLoggingMiddleware(logger log.Logger, s service.Service) service.Service {
	return &loggingMiddleware{logger, s}
}

func (mw loggingMiddleware) Create(restaurant domain.Restaurant) (string, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Create", "took", time.Since(begin))
	}(time.Now())
	return mw.next.Create(restaurant)
}

func (mw loggingMiddleware) Get(id string) (domain.Restaurant, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Get", "took", time.Since(begin))
	}(time.Now())
	return mw.next.Get(id)
}

func (mw loggingMiddleware) Update(id string, restaurant domain.Restaurant) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Update", "took", time.Since(begin))
	}(time.Now())
	return mw.next.Update(id, restaurant)
}

func (mw loggingMiddleware) Delete(id string) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Delete", "took", time.Since(begin))
	}(time.Now())
	return mw.next.Delete(id)
}
