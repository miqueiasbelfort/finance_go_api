package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Mongo_uri = ""
	Port      = 0
)

func Env() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 5000
	}

	Mongo_uri = os.Getenv("MONGO_URI")
}
