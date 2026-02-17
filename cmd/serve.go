package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	prdctHandler "ecommerce/rest/handlers/product"
	usrHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {

	//config
	cnf := config.GetConfig()

	//database config
	dbCon, err := db.NewConnection(cnf.DBConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migration")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//repo
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domain
	user := user.NewService(userRepo)
	product := product.NewService(productRepo)

	//middleware
	middlewares := middleware.NewMiddlewares(cnf)

	//handler
	productHandler := prdctHandler.NewHandler(middlewares, product)
	userHandler := usrHandler.NewHandler(cnf, user)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
