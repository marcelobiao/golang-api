package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8000", r))
}
