package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Start() {

	cnf := config.GetConfig()

	mux := http.NewServeMux()
	mngr := middleware.NewManager()
	mngr.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	InitRoutes(mux, mngr)
	wrappedMux := mngr.WrapMux(mux)

	port := strconv.Itoa(cnf.HttpPort)
	
	fmt.Println("Server running on port", port)
	err := http.ListenAndServe(":" + port, wrappedMux)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}
}
