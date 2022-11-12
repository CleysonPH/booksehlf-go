package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	port     string
	env      string
	host     string
	dsn      string
	dbdriver string
)

func Init(enviroment string) {
	env = enviroment
	if IsDev() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	port = os.Getenv("PORT")
	host = os.Getenv("HOST")
	dsn = os.Getenv("DSN")
	dbdriver = os.Getenv("DBDRIVER")
}

func Addr() string {
	return host + ":" + port
}

func DSN() string {
	return dsn
}

func DBDriver() string {
	return dbdriver
}

func IsDev() bool {
	return env == "dev"
}
