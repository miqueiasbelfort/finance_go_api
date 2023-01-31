package main

import (
	"api/src/config"
	"api/src/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Env()

	r := routers.Init()

	fmt.Printf("The server is running in port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
