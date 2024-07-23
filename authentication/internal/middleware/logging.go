package middleware

import (
	"authentication/internal/domain"
	"authentication/internal/service"
	"time"

	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.AuthService
}

func NewLoggingMiddleware(logger log.Logger, s service.AuthService) service.AuthService {
	return &loggingMiddleware{logger, s}
}

func (mw loggingMiddleware) Login(username, password string) (string, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Login", "took", time.Since(begin))
	}(time.Now())
	return mw.next.Login(username, password)
}

func (mw loggingMiddleware) Register(username, password string, role domain.Role) (string, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Register", "took", time.Since(begin))
	}(time.Now())
	return mw.next.Register(username, password, role)
}
