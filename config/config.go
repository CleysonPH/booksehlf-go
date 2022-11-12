package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	port string
	env  string
	host string
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
}

func Addr() string {
	return host + ":" + port
}

func IsDev() bool {
	return env == "dev"
}
