package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	//InitRoutes(mux, mngr)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)

	port := strconv.Itoa(server.cnf.HttpPort)

	fmt.Println("Server running on port", port)
	err := http.ListenAndServe(":"+port, wrappedMux)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}
}
