package main

import (
	"api/src/config"
	router "api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Load()

	fmt.Println("Rodando api na pota ", config.Port)

	router := router.Init()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
