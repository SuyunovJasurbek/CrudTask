package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	HTTPPort string
	HTTPHost string

	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPassword string
	PostgresDatabase string

	
}

func getOrReturnDefaultValue(key string , defaultValue interface {}) interface {} {
	val , exists :=os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}

func NewConfig() *Config {
	if err := godotenv.Load() ; err!= nil {
		fmt.Println("No .env file found")
	}

	cfg :=Config{}

	cnf := Config{}
	cnf.HTTPPort = getOrReturnDefaultValue("HTTP_PORT", 8080).(string)
	cnf.HTTPHost = getOrReturnDefaultValue("HTTP_HOST", "localhost").(string)

	cnf.PostgresHost = getOrReturnDefaultValue("POSTGRES_HOST", "localhost").(string)
	cnf.PostgresPort = getOrReturnDefaultValue("POSTGRES_PORT", 5432).(string)
	cnf.PostgresUser = getOrReturnDefaultValue("POSTGRES_USER", "postgres").(string)
	cnf.PostgresDatabase = getOrReturnDefaultValue("POSTGRES_DB", "postgres").(string)
	cnf.PostgresPassword = getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres").(string)
	
	return &cfg
}