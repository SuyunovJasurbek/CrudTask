package main

import (
	"fmt"

	"github.com/SuyunovJasurbek/CrudTask/config"
	"github.com/SuyunovJasurbek/CrudTask/router"
	"github.com/SuyunovJasurbek/CrudTask/src/handler"
	"github.com/SuyunovJasurbek/CrudTask/src/repository"
	"github.com/SuyunovJasurbek/CrudTask/src/repository/postgres"
	"github.com/SuyunovJasurbek/CrudTask/src/service"
)

func main() {
	
	cfg :=config.Load()
	pgdb := postgres.NewPostgresDB(cfg)
	r :=repository.NewRepository(pgdb)
	s :=service.NewService(r)
	h :=handler.NewHandler(s)
	w :=router.SetUpRouter(h)
	w.Run(fmt.Sprintf(":%d", cfg.HTTPPort))
}