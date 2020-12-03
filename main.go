package main

import (
	"log"
	"net/http"
	"os"

	"github.com/p886/go-app/handler"
)

func main() {
	http.Handle("/", handler.HelloHandler{})

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Missing env variable 'PORT'")
	}

	log.Printf("Listening on port %s", port)
	http.ListenAndServe(":"+port, nil)
}
