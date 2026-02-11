package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(
		http.HandlerFunc(
			h.GetProducts,
		),
	))

	mux.Handle("POST /products", mngr.With(
		http.HandlerFunc(
			h.CreateProduct,
		),
		middleware.AuthenticateJWT,
	))

	mux.Handle("GET /products/{id}", mngr.With(
		http.HandlerFunc(
			h.GetProduct,
		),
	))

	mux.Handle("PUT /products/{id}", mngr.With(
		http.HandlerFunc(
			h.UpdateProduct,
		),
		middleware.AuthenticateJWT,
	))

	mux.Handle("DELETE /products/{id}", mngr.With(
		http.HandlerFunc(
			h.DeleteProduct,
		),
		middleware.AuthenticateJWT,
	))

}
