package main

//go:generate sh ./pretest.sh

import (
	"FORMULARIO_WEB_MONGODB_GO/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
