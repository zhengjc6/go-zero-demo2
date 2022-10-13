package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type RouteExampleMiddleware struct {
}

func NewRouteExampleMiddleware() *RouteExampleMiddleware {
	return &RouteExampleMiddleware{}
}

func (m *RouteExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("example middle init")
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Info("example middle")
		// Passthrough to next handler if need
		next(w, r)
	}
}
