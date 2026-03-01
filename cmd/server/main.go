package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"itk-test-task/iternal/config"
	"itk-test-task/iternal/handler"
	"itk-test-task/iternal/repository/postgres"
	"itk-test-task/iternal/service"
	"net/http"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	connector, err := postgres.NewDBConnector(cfg)
	if err != nil {
		panic("cannot connect to db")
	}

	walletRepository := postgres.NewWalletRepository(connector)
	walletService := service.NewWalletService(walletRepository)
	walletHandler := handler.NewHandler(walletService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/v1", walletHandler.NewRouter())

	if err := http.ListenAndServe(":"+cfg.HttpListenAddress, r); err != nil {
		panic(err)
	}
}
