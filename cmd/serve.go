package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mngr := middleware.NewManager()
	mngr.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	InitRoutes(mux, mngr)

	wrappedMux := mngr.WrapMux(mux)
	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}
}
