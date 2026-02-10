package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(
		http.HandlerFunc(
			handlers.GetProducts,
		),
	))

	mux.Handle("POST /products", mngr.With(
		http.HandlerFunc(
			handlers.CreateProduct,
		),
		middleware.AuthenticateJWT,
	))

	mux.Handle("GET /products/{id}", mngr.With(
		http.HandlerFunc(
			handlers.GetProduct,
		),
	))

	mux.Handle("PUT /products/{id}", mngr.With(
		http.HandlerFunc(
			handlers.UpdateProduct,
		),
		middleware.AuthenticateJWT,
	))

	mux.Handle("DELETE /products/{id}", mngr.With(
		http.HandlerFunc(
			handlers.DeleteProduct,
		),
		middleware.AuthenticateJWT,
	))

	mux.Handle("POST /users", mngr.With(
		http.HandlerFunc(
			handlers.CreateUser,
		),
	))

	mux.Handle("POST /users/login", mngr.With(
		http.HandlerFunc(
			handlers.LoginUser,
		),
	))
}
