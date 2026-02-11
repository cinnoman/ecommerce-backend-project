package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("POST /users", mngr.With(
		http.HandlerFunc(
			h.CreateUser,
		),
	))

	mux.Handle("POST /users/login", mngr.With(
		http.HandlerFunc(
			h.LoginUser,
		),
	))
}
