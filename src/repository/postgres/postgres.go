package postgres

import (
	"fmt"
	"log"

	"github.com/SuyunovJasurbek/CrudTask/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)


func NewPostgresDB(cfg config.Config) *sqlx.DB {
	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)
	fmt.Println(psqlConnString)
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	} else {
		fmt.Printf("Connected to database: %s \n", cfg.PostgresDatabase)
	}

	return db
}
