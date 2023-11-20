package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	config.Carregar()
}

func main() {
	r := router.Gerar()
	fmt.Printf("Aplicação rodando http://localhost:%d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
