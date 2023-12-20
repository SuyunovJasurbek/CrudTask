package main

import (
	"fmt"

	"github.com/SuyunovJasurbek/CrudTask/api"
	"github.com/SuyunovJasurbek/CrudTask/config"
	_ "github.com/SuyunovJasurbek/CrudTask/docs"
	"github.com/SuyunovJasurbek/CrudTask/src/handler"
	"github.com/SuyunovJasurbek/CrudTask/src/repository"
	"github.com/SuyunovJasurbek/CrudTask/src/repository/postgres"
	"github.com/SuyunovJasurbek/CrudTask/src/service"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	cfg := config.Load()
	pgdb := postgres.NewPostgresDB(cfg)
	defer pgdb.Close()
	r := repository.NewRepository(pgdb)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	w := api.SetUpApi(h)
	w.Run(fmt.Sprintf(":%d", cfg.HTTPPort))
}
