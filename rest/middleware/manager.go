package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleware []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	return n
}

func (mngr *Manager) WrapMux(next http.Handler) http.Handler {
	n := next

	for _, middleware := range mngr.globalMiddleware {
		n = middleware(n)
	}

	return n
}
