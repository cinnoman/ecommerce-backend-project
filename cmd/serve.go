package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /noman", middleware.Logger(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))
	mux.Handle("POST /products-c", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{id}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))

	globalRouterHandler := global_router.GlobalRouter(mux)
	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", globalRouterHandler)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}
}
